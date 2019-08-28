package grpc

import (
	"timingsystem/timingserver/hubs"
	"context"
	"log"
	"net"
	"sync"

	"google.golang.org/grpc"
	pb "timingsystem/timingserver/sysprotos"

	"timingsystem/timingserver/data"
	"timingsystem/timingserver/cerror"
)

const (
	port = ":50051"
)

// server is used to implement timingsystem.TimingSystemServer.
type timingService struct{
	serverHub *hubs.ServerHub
}

// RecordTimingPoint is a method from protobuf.
// test-client will send dummy data to this service by this method
func (s *timingService) RecordTimingPoint(ctx context.Context, in *pb.TimingSystemRequest) (*pb.TimingSystemResponse, error) {
	log.Printf("Id: %v, Type: %v, TimePoint: %v", in.Id, in.Type, in.TimePoint)

	handleTheRequest(in, s.serverHub)

	return &pb.TimingSystemResponse{ResultStatus: true}, nil
}

// StartServing starts grpc serving
func StartServing(serverHub *hubs.ServerHub, wg *sync.WaitGroup) {
	defer wg.Done()

	log.Println("start grpc service!");
	
	if !data.InitialDatabase() {
		log.Fatalf("Data initialization failed!")
	}

	listen, err := net.Listen("tcp4", port)
	cerror.CheckErr(err)

	grpcServer := grpc.NewServer() 
	pb.RegisterTimingSystemServer(grpcServer, &timingService{ serverHub: serverHub })

	err = grpcServer.Serve(listen)
	cerror.CheckErr(err)

}