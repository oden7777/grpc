# baker API

## generate server
```shell
$ protoc \
    --go_out=./api --go-grpc_out=./api \
    proto/*.proto
```

## run
```shell
$ go run api/server/server.go
```

## more
```shell
# server info
$ grpc_cli ls localhost:50051
# server info(detail)
$ grpc_cli ls localhost:50051 -l
# type info
$ grpc_cli type localhost:50051 pancake.baker.Pancake
# call request
$ grpc_cli call localhost:50051 pancake.baker.PancakeBakerService.Bake 'menu:1'
$ grpc_cli call localhost:50051 pancake.baker.PancakeBakerService.Report ''
```