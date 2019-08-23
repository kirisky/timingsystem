package main

import (
	"timingsystem/cerror"
	"context"
	"log"
	"time"

	pb "timingsystem/sysprotos"

	"google.golang.org/grpc"
)

const (
	address = "127.0.0.1:50051"
)

func sendRequests(client pb.TimingSystemClient) {
	data := getFakeData()

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	log.Println("start sending data now!")

	for _, req := range data {
		time.Sleep(2 * time.Second)
		r, err := client.RecordTimingPoint(ctx, req)
		cerror.CheckErr(err)
		log.Printf("Result of Sending data[%v]: %v", req.Id, r.ResultStatus)
	}

	log.Println("Finish sending data!")
}

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	cerror.CheckErr(err)
	defer conn.Close()

	client := pb.NewTimingSystemClient(conn)

	// ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	// defer cancel()

	log.Println("Start sending requests")
	sendRequests(client)
}
