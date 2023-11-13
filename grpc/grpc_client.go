package grpc

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"trace-app/grpc/pb"
)

func DoGrpcRequest(url string) {
	if url == "" {
		url = "localhost:8080"
	}
	log.Println("grpc request is begin.")
	go func() {
		for {
			opts := grpc.WithInsecure()
			cc, err := grpc.Dial(url, opts)
			if err != nil {
				log.Fatal(err)
			}
			defer cc.Close()

			client := pb.NewGreetingServiceClient(cc)
			request := &pb.GreetingServiceRequest{Name: "Tubiers"}

			resp, err := client.Greeting(context.Background(), request)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Printf("Receive grpc response => %s from => %s \n", resp.Message, url)
		}
	}()

}
