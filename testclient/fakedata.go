package main

import (
	pb "timingsystem/timingserver/sysprotos"
)

type fakedata []*pb.TimingSystemRequest

func getFakeData() fakedata {

	data := fakedata {
		// Finish corridor timepoints
		{ Id: int32(1), Type: pb.TimingSystemRequest_FINISH_CORRIDOR, TimePoint: "2019-08-08 20:00:00" },
		{ Id: int32(2), Type: pb.TimingSystemRequest_FINISH_CORRIDOR, TimePoint: "2019-08-08 20:01:00" },
		{ Id: int32(3), Type: pb.TimingSystemRequest_FINISH_CORRIDOR, TimePoint: "2019-08-08 20:02:00" },
		{ Id: int32(4), Type: pb.TimingSystemRequest_FINISH_CORRIDOR, TimePoint: "2019-08-08 20:03:00" },
		{ Id: int32(5), Type: pb.TimingSystemRequest_FINISH_CORRIDOR, TimePoint: "2019-08-08 20:04:00" },
		{ Id: int32(6), Type: pb.TimingSystemRequest_FINISH_CORRIDOR, TimePoint: "2019-08-08 20:05:00" },
		{ Id: int32(7), Type: pb.TimingSystemRequest_FINISH_CORRIDOR, TimePoint: "2019-08-08 20:06:00" },
		{ Id: int32(8), Type: pb.TimingSystemRequest_FINISH_CORRIDOR, TimePoint: "2019-08-08 20:07:00" },
		{ Id: int32(9), Type: pb.TimingSystemRequest_FINISH_CORRIDOR, TimePoint: "2019-08-08 20:08:00" },
		{ Id: int32(10), Type: pb.TimingSystemRequest_FINISH_CORRIDOR, TimePoint: "2019-08-08 20:09:00" },

		// Finishi line timepoints
		{ Id: int32(1), Type: pb.TimingSystemRequest_FINISH_LINE, TimePoint: "2019-08-08 20:04:01" },
		{ Id: int32(2), Type: pb.TimingSystemRequest_FINISH_LINE, TimePoint: "2019-08-08 20:04:04" },
		{ Id: int32(3), Type: pb.TimingSystemRequest_FINISH_LINE, TimePoint: "2019-08-08 20:03:03" },
		{ Id: int32(4), Type: pb.TimingSystemRequest_FINISH_LINE, TimePoint: "2019-08-08 20:03:10" },
		{ Id: int32(5), Type: pb.TimingSystemRequest_FINISH_LINE, TimePoint: "2019-08-08 20:06:09" },
		{ Id: int32(6), Type: pb.TimingSystemRequest_FINISH_LINE, TimePoint: "2019-08-08 20:07:00" },
		{ Id: int32(7), Type: pb.TimingSystemRequest_FINISH_LINE, TimePoint: "2019-08-08 20:08:20" },
		{ Id: int32(8), Type: pb.TimingSystemRequest_FINISH_LINE, TimePoint: "2019-08-08 20:11:02" },
		{ Id: int32(9), Type: pb.TimingSystemRequest_FINISH_LINE, TimePoint: "2019-08-08 20:09:00" },
		{ Id: int32(10), Type: pb.TimingSystemRequest_FINISH_LINE, TimePoint: "2019-08-08 20:10:00" },
	}

	return data
}