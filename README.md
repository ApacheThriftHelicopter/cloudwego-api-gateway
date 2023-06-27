# CloudWeGo-API-Gateway

## Introduction 
This is an API Gateway made with Go frameworks [Kitex](https://github.com/cloudwego/kitex) and [Hertz](https://github.com/cloudwego/hertz) under CloudWeGo developed by ByteDance. The general design for our API Gateway is with the [Backend For Frontend (BFF)](https://medium.com/mobilepeople/backend-for-frontend-pattern-why-you-need-to-know-it-46f94ce420b0) pattern in mind.

## Gateway Flow
![API Gateway Flow][api-gateway-flow]

[api-gateway-flow]: https://github.com/shawnnygoh/cloudwego-api-gateway/blob/main/images/TikTok%20API%20Gateway%20Design%20Architecture.png "API Gateway Flow"

## Purpose of API Gateway
Our API Gateway allows users to easily add their own services following the Thrift IDL specification. This hassle-free approach enables them to quickly build and implement their own services which can be accessed through the Gateway, which will route requests to the appropriate services.

## Basic Features
- Generic Call with Kitex
- Service Registry and Discovery with Nacos
- Load Balancing with Kitex
- Service Mapping 

### Generic Call
Kitex provides a [JSON Mapping Generic Call](https://www.cloudwego.io/docs/kitex/tutorials/advanced-feature/generic-call/#4-json-mapping-generic-call) feature which helps with Thrift codec. 

### Service Registry & Discovery
A service registry is required to keep track of the available instances of each service. This gateway utilizes a [Kitex extension](https://github.com/kitex-contrib/registry-nacos) for server registration and discovery with Nacos.  

### Load Balancing
Kitex provides a few [load balancer](https://www.cloudwego.io/docs/kitex/tutorials/service-governance/loadbalance/) options to choose from with WeightedRoundRobin, Weighted Random and ConsistentHash. For this gateway, we have opted for the WeightedRoundRobin load balancer which distributes any incoming requests to the backend Kitex servers based on weight.

### Service Mapping
In order to promote maintainability for the API Gateway, we also provide a service mapping feature such that we can not only establish the relationship between IDLs and services, but also ensure that future updates to the services through the IDL files will be reflected in the actual service itself.  

## Documentation

## Set-up
To start working with our API Gateway, users will have to first install Nacos [here](https://nacos.io/en-us/docs/v2/quickstart/quick-start.html). 

Users can add their own services by providing their own code which follows the [Thrift IDL specification](https://thrift.apache.org/docs/idl) and storing them in the `idl` directory. The file should include requests, responses and the service itself and might look something like the following:

```thrift
namespace go book

struct QueryBookReq {
    1: i32 Num (api.query="num", api.vd="$<100; msg:'num must less than 100'");
}

struct QueryBookResp {
    1: string ID;
    2: string Title;
    3: string Author;
    4: string Content; 
}

struct InsertBookReq {
    1: string ID (api.form="id");
    2: string Title (api.form="title");
    3: string Author (api.form="author");
}

struct InsertBookResp {
    1: bool Ok; 
    2: string Msg; 
}

service BookSvc {
   QueryBookResp queryBook(1: QueryBookReq req) (api.get="book/query");
   InsertBookResp insertBook(1: InsertBookReq req) (api.post="book/insert");
}
```

## Performance