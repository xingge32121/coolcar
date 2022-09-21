PROTO_PATH=./rental/api
GO_OUT_PATH=./rental/api/gen/v1

protoc -I=$PROTO_PATH --go_out=paths=source_relative:$GO_OUT_PATH rental.proto
protoc -I=$PROTO_PATH --go-grpc_out=paths=source_relative:$GO_OUT_PATH rental.proto
protoc -I=$PROTO_PATH --grpc-gateway_out=paths=source_relative,grpc_api_configuration=$PROTO_PATH/rental.yaml:$GO_OUT_PATH rental.proto