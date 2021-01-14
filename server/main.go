package main

import (
	"log"
	"net"
	user "viniciuslimafernandes/go-grpc-profile-info/proto/gen"

	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp",":8182")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	usr := Server{}

	grpcServer := grpc.NewServer()

	user.RegisterUserServiceServer(grpcServer, &usr)

	log.Println("Listening on port: 8182")

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %s", err)
	}
}