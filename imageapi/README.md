# Image Upload API

## generate server
```shell
$ protoc \
    --go_out=./api --go-grpc_out=./api \
    proto/*.proto
```

## run
```shell
# run server
$ go run api/server/server.go
# run example client
$ go run cmd/client.go /path/to/filepath
```

## more
```shell
# server info
$ grpc_cli ls localhost:50051
# server info(detail)
$ grpc_cli ls localhost:50051 -l
# type info
$ grpc_cli type localhost:50051 image.uploader.ImageUploadRequest
# call request
# $ grpc_cli call localhost:50051 image.uploader.ImageUploadService.Upload 'file_meta:{filename:""}'
```