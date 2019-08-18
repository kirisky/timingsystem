package main

import (
	"context"
	"log"
	"net"

	pb "automatic-timing-system/sysprotos"

	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

func initialDatabase() (result bool) {
	log.Println("this is a function for initialing database")
	return true
}

// server is used to implement timingsystem.TimingSystemServer.
type timingService struct{}

func (s *timingService) RecordTimingPoint(ctx context.Context, in *pb.TimingSystemRequest) (*pb.TimingSystemResponse, error) {
	log.Printf("Id: %v, Type: %v, TimePoint: %v", in.Id, in.Type, in.TimePoint)
	TestIsAHaha()
	return &pb.TimingSystemResponse{ResultStatus: true}, nil
}

func main() {
	initialDatabase()

	listen, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterTimingSystemServer(s, &timingService{})

	if err := s.Serve(listen); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
