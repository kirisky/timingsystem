package main

import (
	"context"
	"log"
	"time"

	pb "timingsystem/sysprotos"

	"google.golang.org/grpc"
)

const (
	address = "localhost:50051"
)

func getRequest() (request *pb.TimingSystemRequest) {
	return &pb.TimingSystemRequest{
		Id:        100,
		Type:      pb.TimingSystemRequest_FINISH_CORRIDOR,
		TimePoint: "2019-08-08 20:00:00",
	}
}

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure())

	if err != nil {
		log.Fatalf("cannot connect to the server: %v", err)
	}

	defer conn.Close()

	client := pb.NewTimingSystemClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	request := getRequest()
	r, err := client.RecordTimingPoint(ctx, request)

	if err != nil {
		log.Fatalf("Send data failed: %v", err)
	}
	log.Printf("Result of Sending data: %v", r.ResultStatus)
}
