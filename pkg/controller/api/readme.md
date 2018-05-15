# Create gRPC go files
protoc -I$GOPATH/src -I. --gofast_out=plugins=grpc:. controller.proto
