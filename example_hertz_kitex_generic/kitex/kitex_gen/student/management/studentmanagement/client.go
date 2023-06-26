// Code generated by Kitex v0.5.2. DO NOT EDIT.

package studentmanagement

import (
	"context"
	management "example_hertz_kitex/kitex/kitex_gen/student/management"
	client "github.com/cloudwego/kitex/client"
	callopt "github.com/cloudwego/kitex/client/callopt"
)

// Client is designed to provide IDL-compatible methods with call-option parameter for kitex framework.
type Client interface {
	QueryStudent(ctx context.Context, req *management.QueryStudentRequest, callOptions ...callopt.Option) (r *management.QueryStudentResponse, err error)
	InsertStudent(ctx context.Context, req *management.InsertStudentRequest, callOptions ...callopt.Option) (r *management.InsertStudentResponse, err error)
}

// NewClient creates a client for the service defined in IDL.
func NewClient(destService string, opts ...client.Option) (Client, error) {
	var options []client.Option
	options = append(options, client.WithDestService(destService))

	options = append(options, opts...)

	kc, err := client.NewClient(serviceInfo(), options...)
	if err != nil {
		return nil, err
	}
	return &kStudentManagementClient{
		kClient: newServiceClient(kc),
	}, nil
}

// MustNewClient creates a client for the service defined in IDL. It panics if any error occurs.
func MustNewClient(destService string, opts ...client.Option) Client {
	kc, err := NewClient(destService, opts...)
	if err != nil {
		panic(err)
	}
	return kc
}

type kStudentManagementClient struct {
	*kClient
}

func (p *kStudentManagementClient) QueryStudent(ctx context.Context, req *management.QueryStudentRequest, callOptions ...callopt.Option) (r *management.QueryStudentResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.QueryStudent(ctx, req)
}

func (p *kStudentManagementClient) InsertStudent(ctx context.Context, req *management.InsertStudentRequest, callOptions ...callopt.Option) (r *management.InsertStudentResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.InsertStudent(ctx, req)
}
