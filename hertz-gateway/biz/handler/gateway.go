package handler

import (
	"context"
	"net/http"
	"encoding/json"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"

	"github.com/cloudwego/kitex/pkg/kerrors"

	errors "github.com/ApacheThriftHelicopter/cloudwego-api-gateway/hertz-gateway/biz/errors"
	servicemapping "github.com/ApacheThriftHelicopter/cloudwego-api-gateway/hertz-gateway/biz/servicemapping"
)

// Required parameters of the user request
type requiredParams struct {
	BizParams string `json:"biz_params"`
}

// Service Map used to store IDL mappings
var SvcMap = servicemapping.ServiceMap{}

// Gateway handle the request with the query path of prefix `/gateway`.
func Gateway(ctx context.Context, c *app.RequestContext) {
	// Look for svc and method parameters
	svcName := c.Param("svc")
	methodName := c.Param("method")
	hlog.Info(methodName)

	cli, ok := SvcMap.Map[svcName]
	if !ok {
		c.JSON(http.StatusOK, errors.New(errors.Err_BadRequest))
		return
	}
	hlog.Info(cli)

	var params requiredParams
	hlog.Info(params)
	if err := c.BindAndValidate(&params); err != nil {
		hlog.Error(err)
		c.JSON(http.StatusOK, errors.New(errors.Err_ServerMethodNotFound))
		return
	}
	resp, err := cli.GenericCall(ctx, methodName, params.BizParams)
	hlog.Info(resp)
	respMap := make(map[string]interface{})
	if err != nil {
		hlog.Errorf("GenericCall err: %v", err)
		bizErr, ok := kerrors.FromBizStatusError(err)
		if !ok {
			c.JSON(http.StatusOK, errors.New(errors.Err_ServerHandleFail))
			return
		}
		respMap["err_code"] = bizErr.BizStatusCode()
		respMap["err_message"] = bizErr.BizMessage()
		c.JSON(http.StatusOK, respMap)
		return
	}
	realResp, ok := resp.([]byte)
	if !ok {
		c.JSON(http.StatusOK, errors.New(errors.Err_ServerHandleFail))
		return
	}
	JSONResp := make(map[string]string)
	if err := json.Unmarshal(realResp, &JSONResp); err != nil {
		c.JSON(http.StatusOK, errors.New(errors.Err_ResponseUnableParse))
	}
	JSONResp["err_code"] = "0"
	JSONResp["err_msg"] = "ok"
	c.JSON(http.StatusOK, JSONResp)
}