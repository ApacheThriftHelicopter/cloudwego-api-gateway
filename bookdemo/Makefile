kitex_gen_user:
	kitex --thrift-plugin validator -module github.com/ApacheThriftHelicopter/cloudwego-api-gateway/bookdemo idl/user.thrift # execute in the project root directory

kitex_gen_book:
	kitex --thrift-plugin validator -module github.com/ApacheThriftHelicopter/cloudwego-api-gateway/bookdemo idl/book.thrift # execute in the project root directory

install_hz_latest:
	go install github.com/cloudwego/hertz/cmd/hz@latest


hertz_gen_model:
	hz model --idl=idl/bookapi.thrift --mod=github.com/ApacheThriftHelicopter/cloudwego-api-gateway/bookdemo --model_dir=hertz_gen

hertz_gen_client:
	hz client --idl=idl/bookapi.thrift --base_domain=http://127.0.0.1:8080 --client_dir=api_request --mod=github.com/ApacheThriftHelicopter/cloudwego-api-gateway/bookdemo --model_dir=hertz_gen