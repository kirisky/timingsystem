package main

import (
	"context"
	"log"
	"time"
	"flag"

	pb "timingsystem/testclient/timingsystem"

	"google.golang.org/grpc"
)

const (
	// address = "127.0.0.1:50051"
	// address = "35.228.42.74:50051"
)

var addr = flag.String("addr", ":50051", "gRPC service address")

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
	flag.Parse()

	conn, err := grpc.Dial(*addr, grpc.WithInsecure())
	if err != nil {
		log.Fatalln("Dial gRPC is failed: ", err)
	}
	defer conn.Close()

	client := pb.NewTimingSystemClient(conn)

	log.Println("Start sending requests")
	sendRequests(client)
}
