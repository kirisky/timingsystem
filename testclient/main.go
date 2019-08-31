package main

import (
	"context"
	"log"
	"time"

	pb "timingsystem/testclient/timingsystem"

	"google.golang.org/grpc"
)

const (
	address = "127.0.0.1:50051"
)

func sendRequests(client pb.TimingSystemClient) {
	data := getFakeData()
	seconds := time.Duration(len(data) * 3)

	ctx, cancel := context.WithTimeout(context.Background(), seconds * time.Second)
	defer cancel()

	log.Println("start sending data now!")

	for _, req := range data {
		time.Sleep(2 * time.Second)
		log.Printf("req: [%v]", req)
		r, err := client.RecordTimingPoint(ctx, req)
		if err != nil {
			log.Fatalln("Calling RecordTimePoint is failed: ", err)
		}
		log.Printf("Result of Sending data[%v]: %v", req.Id, r.ResultStatus)
	}

	log.Println("Finish sending data!")
}

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalln("Dial gRPC is failed: ", err)
	}
	defer conn.Close()

	client := pb.NewTimingSystemClient(conn)

	// ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	// defer cancel()

	log.Println("Start sending requests")
	sendRequests(client)
}
