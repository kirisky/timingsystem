package grpc

import (
	"log"
	
	pb "timingsystem/sysprotos"

	"timingsystem/data"
)

func handleTheRequest(in *pb.TimingSystemRequest) {

	athletes := data.GetRecords(int(in.Id), in.Type, in.TimePoint)
	if len(athletes) == 0 {
		log.Fatal("Get records failed!")
	}

	for key, val := range athletes {
		log.Println(key, val.ID, val.FullName, val.StartNumber, val.FinishCorridorTime, val.FinishLineTime)	
	}
}