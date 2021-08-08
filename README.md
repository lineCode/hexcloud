![Build and Deploy to Cloud Run](https://github.com/3vilM33pl3/hexcloud/workflows/Build%20and%20Deploy%20to%20Cloud%20Run/badge.svg)

## Build
```bash
docker build -f ./deploy/hexcloud/Dockerfile -t hexcloud .
```

## Update gRPC proto
```shell
go install google.golang.org/protobuf/cmd/protoc-gen-go
install google.golang.org/grpc/cmd/protoc-gen-go-grpc
protoc --go_out=./internal/pkg  --go-grpc_out=./internal/pkg ./api/hexagon.proto
```
