package main

import (
	"context"
	"log"
	"net"

	"proto" // Импортируем сгенерированные файлы

	"google.golang.org/grpc"
)

type server struct{}

func (s *server) Add(ctx context.Context, in *proto.AddRequest) (*proto.AddResponse, error) {
	result := in.A + in.B
	return &proto.AddResponse{Result: result}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	proto.RegisterCalculatorServer(s, &server{})

	log.Println("Server started at :50051")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
