package main

import (
	pb "github.com/mehrdadrad/GoBestPractices/grpc/routeguide"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"log"
	"math/rand"
	"net"
)

const (
	port = ":55555"
)

type server struct{}

func (s *server) GetRand(ctx context.Context, in *pb.RandRequest) (*pb.RandResponse, error) {
	return &pb.RandResponse{RandNum: rand.Int63n(in.MaxNum)}, nil
}

func gRPCServer() {
	l, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatal(err)
	}
	s := grpc.NewServer()
	pb.RegisterJaguarServer(s, &server{})
	s.Serve(l)
}

func main() {
	gRPCServer()
}
