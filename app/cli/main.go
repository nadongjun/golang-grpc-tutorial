package main

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"time"

	pb "go_project/service"

	"google.golang.org/grpc"
)

const (
	address = "localhost:50051"
)

var intStream chan int

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
	log.Println(r.GetId())
	id, err := strconv.Atoi(r.GetId())
	intStream <- id
	defer close(intStream)
	defer fmt.Println("Grpc Done.")
}

func main() {
	intStream = make(chan int, 4)
	go InsertData("Hi")

	for intValue := range intStream {
		fmt.Println("Insert Data Id :", intValue)
	}

}
