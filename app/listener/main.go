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

func GetData(id string) {
	log.Println("InsertData")
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewServiceInterfaceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.GetData(ctx, &pb.GetMsg{Id: id})
	if err != nil {
		log.Fatalf("%v", err)
	}
	log.Printf("Data : %s", r.GetMsg())

}

func main() {
	GetData("Hi")
}
