package grpc

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"
	pb "timingsystem/sysprotos"

	"timingsystem/data"
	"timingsystem/cerror"
)

const (
	port = ":50051"
	// host = "127.0.0.1"
	// host = "0.0.0.0"
)

// server is used to implement timingsystem.TimingSystemServer.
type timingService struct{}

func (s *timingService) RecordTimingPoint(ctx context.Context, in *pb.TimingSystemRequest) (*pb.TimingSystemResponse, error) {
	log.Printf("Id: %v, Type: %v, TimePoint: %v", in.Id, in.Type, in.TimePoint)
	
	handleTheRequest(in)

	return &pb.TimingSystemResponse{ResultStatus: true}, nil
}

// StartServing starts grpc serving
func StartServing() {
	log.Println("start serving!");
	
	if !data.InitialDatabase() {
		log.Fatalf("Data initialization failed!")
	}

	listen, err := net.Listen("tcp4", port)
	cerror.CheckErr(err)

	grpcServer := grpc.NewServer()
	pb.RegisterTimingSystemServer(grpcServer, &timingService{})

	err = grpcServer.Serve(listen)
	cerror.CheckErr(err)

}