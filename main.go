package main

import (
	"flag"
	"fmt"
	"trace-app/grpc"
	"trace-app/server"
)

func main() {
	shutdown := make(chan struct{}, 1)
	grpcUrl := flag.String("grpc", fmt.Sprintf("localhost:8080"), "grpc server url,default localhost:8080")
	httpUrl := flag.String("http", fmt.Sprintf("localhost:9090"), "http url, default localhost:9090")

	go grpc.RunGrpcServer()
	go server.RunRegionZoneServer()
	go server.DoHTTPSRequest()
	go server.DoHTTPRequest(*httpUrl)
	go grpc.DoGrpcRequest(*grpcUrl)
	<-shutdown
}
