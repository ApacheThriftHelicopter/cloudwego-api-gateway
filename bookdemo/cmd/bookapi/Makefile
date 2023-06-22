install_hz_latest:
	go install github.com/cloudwego/hertz/cmd/hz@latest

hertz_new_api:
	hz new --idl=../../idl/bookapi.thrift --service=hello.api --router_dir=hertz_router --handler_dir=hertz_handler -f -use=github.com/ApacheThriftHelicopter/cloudwego-api-gateway/bookdemo/hertz_gen

hertz_update_api:
	hz update --idl=../../idl/bookapi.thrift -use=github.com/ApacheThriftHelicopter/cloudwego-api-gateway/bookdemo/hertz_gen

run:
	sh build.sh && sh output/bootstrap.sh