package grpc

import (
	"timingsystem/timingserver/cerror"
	"log"
	"encoding/json"
	
	pb "timingsystem/sysprotos"
	"timingsystem/timingserver/hubs"

	"timingsystem/timingserver/data"
)

// handleTheRequest gets the currect records and sends the records to websocket service
func handleTheRequest(in *pb.TimingSystemRequest, serverHub *hubs.ServerHub) {

	athletes := data.GetRecords(int(in.Id), in.Type, in.TimePoint)
	if len(athletes) == 0 {
		log.Fatal("Get records failed!")
	}
	log.Println("Athletes Size:", len(athletes))

	athletesJSON, err := json.Marshal(athletes)
	cerror.CheckErr(err)

	serverHub.MessagePipe <- athletesJSON
}