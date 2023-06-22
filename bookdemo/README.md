# BookDemo

## Introduction

This bookdemo is built with CloudWeGo frameworks [Kitex](https://github.com/cloudwego/kitex) and [Hertz](https://github.com/cloudwego/hertz). 


| Service Name | Path         | Corresponding IDL     |
|--------------|--------------|-----------------------|
| bookapi      | cmd/bookapi  | idl/bookapi.thrift    |
| user         | cmd/user     | idl/user.thrift       |
| book         | cmd/book     | idl/book.thrift       |

### Basic Features
- Server Discovery and Registry with `registry-etcd`
- Load Balancing

### Organization 
| folder         | contents                |
|----------------|-------------------------|
| hertz_handler  | HTTP handler            |
| service        | Business logic          |
| rpc            | RPC call logic          |
| dal            | DB operation            |
| pack           | Data pack               |
| pkg/middleware | RPC middleware          |
| pkg/constants  | Constants               |
| pkg/errno      | Customized error number |
| pkg/configs    | SQL and Tracing configs |

## Set-up

```shell
docker-compose up
```

### Start running User RPC Server

```shell
cd cmd/user
sh build.sh
sh output/bootstrap.sh
```

### Start running Book RPC Server

```shell
cd cmd/book
sh build.sh
sh output/bootstrap.sh
```

### Start running BookAPI Server

```shell
cd cmd/bookapi
sh build.sh
sh output/bootstrap.sh
```

## Documentation
1. Create IDLs for services and api which is stored under idl directory.
2. Generate Kitex client scaffolding code with `kitex idl/book.thrift` and `kitex idl/user.thrift`.
3. Create a new directory called cmd and generate Hertz scaffolding code with `hz new --idl=../idl/bookapi.thrift`.
4. Store `model/bookapi.go` in a new directory hertz_gen which will be in the root directory. 
5. Add custom packages (configs, constants, errno and middleware) under the pkg directory.

### cmd directory 
1. Create new directories book, bookapi and user under cmd to store all the hertz generated code and additional code to be written. 

#### bookapi directory 
1. Move previously hertz generated code for bookapi.thrift under the bookapi directory.
2. Rename handler and router to hertz_handler and hertz_router respectively.
3. Create a directory `middleware` for jwt middleware and add middleware in `jwt.go`.
4. Create a directory `rpc` for rpc and add rpc code in `book.go`, `init.go` and `user.go`. 
5. Edit code in `router.go`.
6. Add code to `main.go`.

#### book directory

#### user directory 

### Final steps
1. `go mod init` and `go mod tidy` to import all necessary dependencies. 
2. Add licenses
3. Create `Makefile` and `docker-compose.yaml` for easy set-up.




