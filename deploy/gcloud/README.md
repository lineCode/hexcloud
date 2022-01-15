# Deploy into GCP

## Continuous integration
Github Actions builds and deploys `hexcloud` into GCP Cloud Run

## API Gateway
The API gateway is made separately with a api config and an api descriptor.

[API Gateway and gRPC in GCP](https://cloud.google.com/api-gateway/docs/get-started-cloud-run-grpc)

### API Descriptor
The content of a `.proto` file can be represented using protocol buffers [ref](https://developers.google.com/protocol-buffers/docs/techniques#self-description).
The API Gateway uses this to get the de
Create with the protocol buffer command:

```shell
protoc --go_out=./internal/pkg  --descriptor_set_out=deploy/gcloud/api_descriptor.pb --go-grpc_out=./internal/pkg ./api/hexagon.proto
```


### Service Configuration
The configuration contains the service name and the endpoints
```yaml
type: google.api.Service
config_version: 3
name: "*.apigateway.robot_motel.cloud.goog"
title: API Gateway Cloud Run gRPC gcloud
apis:
  - name: endpoints.hexworld.gcloud.HexagonService
usage:
  rules:
    - selector: endpoints.hexworld.gcloud.HexagonService.GetHexagonRing
      allow_unregistered_calls: true
    - selector: endpoints.hexworld.gcloud.HexagonService.StoreHexagons
      allow_unregistered_calls: true
    - selector: endpoints.hexworld.gcloud.HexagonService.GetStatus
      allow_unregistered_calls: true
backend:
  rules:
    - selector: "*"
      address: grpcs://gcloud-j6feiuh7aa-ue.a.run.app
```

## Manual deploy

```shell
gcloud api-gateway api-configs create hexcloud-config --api=hexcloud-prod --project=robot-motel --grpc-files=api_descriptor.pb,api_config.yaml

gcloud api-gateway gateways create hexcloud-grpc --api=hexcloud-prod --api-config=hexcloud-config --location=us-east1 --project=robot-motel
```

