package main

import (
	"context"
	"log"
	"net"
	"os"

	pb "timingsystem/sysprotos"

	"google.golang.org/grpc"

	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

const (
	port = ":50051"
)

type athleteInfo struct {
	ID          int
	StartNumber int
}

var finishCorridorContainer = make([]athleteInfo, 0)
var finishLineContainer = make([]athleteInfo, 0)
var athletes = make([]athleteInfo, 0)

func initialAthletes() {

	athletes = append(athletes, athleteInfo{ID: 1, StartNumber: 100})
	athletes = append(athletes, athleteInfo{ID: 2, StartNumber: 101})
	athletes = append(athletes, athleteInfo{ID: 3, StartNumber: 102})
	athletes = append(athletes, athleteInfo{ID: 4, StartNumber: 103})
	athletes = append(athletes, athleteInfo{ID: 5, StartNumber: 104})
	athletes = append(athletes, athleteInfo{ID: 6, StartNumber: 105})
	athletes = append(athletes, athleteInfo{ID: 7, StartNumber: 106})
	athletes = append(athletes, athleteInfo{ID: 8, StartNumber: 107})
	athletes = append(athletes, athleteInfo{ID: 9, StartNumber: 108})
	athletes = append(athletes, athleteInfo{ID: 10, StartNumber: 109})
}

func initialDatabase() (result bool) {
	os.Remove("./atheletes.db")

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

	initialAthletes()
	for key, value := range athletes {
		_, err = stmt.Exec(value.ID, key, value.StartNumber)
		checkErr(err)
	}

	return true
}

func putDataIntoContainers(in *pb.TimingSystemRequest) {
	
}


// server is used to implement timingsystem.TimingSystemServer.
type timingService struct{}

func (s *timingService) RecordTimingPoint(ctx context.Context, in *pb.TimingSystemRequest) (*pb.TimingSystemResponse, error) {
	log.Printf("Id: %v, Type: %v, TimePoint: %v", in.Id, in.Type, in.TimePoint)
	return &pb.TimingSystemResponse{ResultStatus: true}, nil
}

func main() {
	// initialDatabase()

	listen, err := net.Listen("tcp", port)
	checkErr(err)

	s := grpc.NewServer()
	pb.RegisterTimingSystemServer(s, &timingService{})

	err = s.Serve(listen)
	checkErr(err)
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
