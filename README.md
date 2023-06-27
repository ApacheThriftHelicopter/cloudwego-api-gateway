# CloudWeGo-API-Gateway
This is an API Gateway made with Go frameworks [Kitex](https://github.com/cloudwego/kitex) and [Hertz](https://github.com/cloudwego/hertz) under CloudWeGo developed by ByteDance. 

## Gateway Flow
![API Gateway Flow][api-gateway-flow]

[api-gateway-flow]: https://github.com/shawnnygoh/cloudwego-api-gateway/blob/main/images/TikTok%20API%20Gateway%20Design%20Architecture.png "API Gateway Flow"

## Purpose of API Gateway
Our API Gateway allows users to easily add their own services following the Thrift IDL specification. This hassle-free approach enables them to quickly build and implement their own services which can be accessed through the Gateway, which will route requests to the appropriate services.

## Basic Features
- Generic Call with Kitex
- Service Registry with Nacos
- Load Balancing with Kitex
- Service Mapping

### Generic Call
[Kitex](https://github.com/cloudwego/kitex) provides a [JSON Mapping Generic Call](https://www.cloudwego.io/docs/kitex/tutorials/advanced-feature/generic-call/#4-json-mapping-generic-call) feature which helps with Thrift codec. 

### Service Registry

### Load Balancing

### Service Mapping

## Documentation

## Set-up
