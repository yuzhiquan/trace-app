package main

import (
	"flag"
	"fmt"
	"trace-app/grpc"
	"trace-app/server"
)

func main() {
	cpc := false
	shutdown := make(chan struct{}, 1)
	grpcUrl := flag.String("grpc", fmt.Sprintf("localhost:8080"), "grpc server url,default localhost:8080")
	httpUrl := flag.String("http", "http://httpbin.org", "http url, default http://httpbin.org")
	httpsUrl := flag.String("https", "https://httpbin.org", "http url, default https://httpbin.org")
	interval := flag.Int("interval", 5, "request interval, default 5")
	flag.BoolVar(&cpc, "cpc", false, "egress under cpc,default is false")
	flag.Parse()
	go grpc.RunGrpcServer()
	go server.RunRegionZoneServer()
	go server.DoHTTPSRequest(*httpsUrl, *interval)
	go server.DoHTTPRequest(*httpUrl, *interval)
	go grpc.DoGrpcRequest(*grpcUrl, *interval, cpc)
	<-shutdown
}
