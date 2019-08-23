package data

import (
	"testing"
	pb "timingsystem/sysprotos"
)

func TestGetRecords(t *testing.T) {
	records := GetRecords(
		1,
		pb.TimingSystemRequest_FINISH_CORRIDOR,
		"2019-08-08 10:20:20",
	)

	if len(records) == 0 {
		t.Error("Get records failed")
	}

	record := records[0]
	if record.ID != 1 {
		t.Error("Get a wrong record")
	}

	t.Log("the record is correct! ")
}