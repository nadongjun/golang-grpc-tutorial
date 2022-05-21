package main

import (
	"context"
	"log"
	"net"

	pb "go_project/service"

	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedServiceInterfaceServer
}

func (s *server) InsertData(ctx context.Context, in *pb.InsertMsg) (*pb.InsertResponse, error) {
	log.Printf("Received Message: %s", in.GetMsg())
	return &pb.InsertResponse{Id: "999", ResponseCode: 200}, nil
}

func (s *server) GetData(ctx context.Context, in *pb.GetMsg) (*pb.GetResponse, error) {
	log.Printf("Received Message: %s", in.GetId())
	return &pb.GetResponse{Msg: "msg", ResponseCode: 200}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterServiceInterfaceServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
