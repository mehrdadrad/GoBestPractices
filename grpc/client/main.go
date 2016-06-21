package main

import (
	pb "github.com/mehrdadrad/GoBestPractices/grpc/routeguide"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"log"
)

const (
	address = "localhost:55555"
)

type server struct{}

func gRPCClient() {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	c := pb.NewJaguarClient(conn)
	r, err := c.GetRand(context.Background(), &pb.RandRequest{MaxNum: 100})
	if err != nil {
		log.Fatal(err)
	}
	println(r.RandNum)
}

func main() {
	gRPCClient()
}
