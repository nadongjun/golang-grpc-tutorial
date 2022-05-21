package main

import (
	"context"
	"log"
	"time"

	pb "go_project/service"

	"google.golang.org/grpc"
)

const (
	address = "localhost:50051"
)

func InsertData(msg string) {
	log.Println("InsertData")
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewServiceInterfaceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.InsertData(ctx, &pb.InsertMsg{Msg: msg})
	if err != nil {
		log.Fatalf("%v", err)
	}
	log.Printf("Data Id : %s", r.GetId())

}

func main() {
	InsertData("Hi")
}
