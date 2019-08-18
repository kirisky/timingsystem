package main

import (
	"context"
	"log"
	"net"

	pb "automatic-timing-system/sysprotos"

	"google.golang.org/grpc"

	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

const (
	port = ":50051"
)

type athleteInfo struct {
	Id int,
	StartNumber int
}

func getAthletes() (athletes map){
	athletes := make(map[string]*atheleteInfo)
	athletes["AthleteA"] = &atheleteInfo { Id = 1; StartNumber = 100 }
	athletes["AthleteB"] = &atheleteInfo { Id = 2; StartNumber = 101 }
	athletes["AthleteC"] = &atheleteInfo { Id = 3; StartNumber = 102 }
	athletes["AthleteD"] = &atheleteInfo { Id = 4; StartNumber = 103 }
	athletes["AthleteE"] = &atheleteInfo { Id = 5; StartNumber = 104 }
	athletes["AthleteF"] = &atheleteInfo { Id = 6; StartNumber = 105 }
	athletes["AthleteG"] = &atheleteInfo { Id = 7; StartNumber = 106 }
	athletes["AthleteH"] = &atheleteInfo { Id = 8; StartNumber = 107 }
	athletes["AthleteI"] = &atheleteInfo { Id = 9; StartNumber = 108 }
	athletes["AthleteG"] = &atheleteInfo { Id = 10; StartNumber = 109 }
}

func initialDatabase() (result bool) {
	os.Remove("sqlite3", "./atheletes.db")

	db, err := sql.Open("sqlite3", "./athletes.db")
	checkErr(err)
	defer db.Close()

	sqlStmt := `
	create table athletes (
		Id integer not null primary key, 
		FullName text, 
		StartNumber integer
	);
	delete from athletes;
	`
	_, err = db.Exec(sqlStmt)
	checkErr(err)

	insertSQL := "INSERT INTO athletes(Id, FullName, StartNumber) values(?, ?, ?)"
	stmt, err := db.Prepare(insertSQL)
	checkErr(err)
	defer stmt.Close()

	for key, value := range athletes {
		_, err = stmt.Exec(value.Id, key, value.StartNumber)	
		checkErr(err)
	}

	return true
}


// server is used to implement timingsystem.TimingSystemServer.
type timingService struct{}

func (s *timingService) RecordTimingPoint(ctx context.Context, in *pb.TimingSystemRequest) (*pb.TimingSystemResponse, error) {
	log.Printf("Id: %v, Type: %v, TimePoint: %v", in.Id, in.Type, in.TimePoint)
	TestIsAHaha()
	return &pb.TimingSystemResponse{ResultStatus: true}, nil
}

func main() {
	// initialDatabase()

	listen, err := net.Listen("tcp", port)
	checkErr(err)

	s := grpc.NewServer()
	pb.RegisterTimingSystemServer(s, &timingService{})

	err := s.Serve(listen);
	checkErr(err)
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
