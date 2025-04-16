package main

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	pb "path/to/your/protobuf/package"
)

const (
	port = ":50051"
)

// server is used to implement the gRPC server.
type server struct {
	pb.UnimplementedYourServiceServer
}

// YourMethod is an example method implementation.
func (s *server) YourMethod(ctx context.Context, in *pb.YourRequest) (*pb.YourResponse, error) {
	// Implement your method logic here
	return &pb.YourResponse{}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterYourServiceServer(s, &server{})
	reflection.Register(s)

	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
