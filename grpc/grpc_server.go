package grpc

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"log"
	"math/rand"
	"net"
	"trace-app/grpc/pb"
)

type server struct {
	pb.GreetingServiceServer
}

var invalidArgument = status.New(codes.InvalidArgument, "invalid args")
var unavailable = status.New(codes.Unavailable, "Unavailable")
var unimplemented = status.New(codes.Unimplemented, "Unimplemented")
var notFound = status.New(codes.NotFound, "NotFound")
var statusList = []status.Status{*invalidArgument, *unavailable, *unimplemented, *notFound}
var ranInt = 0
var randomFail = false

func (s *server) Greeting(ctx context.Context, req *pb.GreetingServiceRequest) (*pb.GreetingServiceReply, error) {
	md, _ := metadata.FromIncomingContext(ctx)
	fmt.Printf("metadata:%+v", md)
	if randomFail && rand.Intn(100)%1 == 0 {
		i := rand.Intn(len(statusList))
		header := metadata.Pairs("grpc-status", statusList[i].Code().String())
		err := grpc.SendHeader(ctx, header)
		if err != nil {
			fmt.Printf("header set failed\n")
		} else {
			fmt.Printf("header set succ\n")
		}
		// create and set trailer
		trailer := metadata.Pairs("grpc-status", statusList[i].Code().String())
		err = grpc.SetTrailer(ctx, trailer)

		if err != nil {
			fmt.Printf("trailer set failed\n")
		} else {
			fmt.Printf("trailer set succ\n")

		}
		//
		//header := metadata.New(map[string]string{"grpc-status": statusList[i].Code().String()})
		//if err := grpc.SendHeader(ctx, header); err != nil {
		//	return nil, status.Errorf(codes.Internal, "unable to send 'grpc-status' header")
		//}
		return nil, statusList[i].Err()
	}

	return &pb.GreetingServiceReply{
		Message: fmt.Sprintf("Hello, %s", req.Name),
	}, nil
}

func RunGrpcServer(random bool) {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}

	if random {
		randomFail = true
	}

	s := grpc.NewServer()
	pb.RegisterGreetingServiceServer(s, &server{})
	if err := s.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
