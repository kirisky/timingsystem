package data

import (
	"timingsystem/timingserver/cerror"
	"log"

	pb "timingsystem/sysprotos"
	"timingsystem/timingserver/utils"

	"database/sql"
	// import go-sqlite3 driver
	_ "github.com/mattn/go-sqlite3"

	"github.com/emirpasic/gods/sets/treeset"
	// "github.com/deckarep/golang-set"
)


func byTimePoint(a, b interface{}) int {
	athleteA := a.(athlete)
	athleteB := b.(athlete)

	if !utils.IsEmptyString(athleteA.FinishLineTime) {
		if utils.IsEmptyString(athleteB.FinishLineTime) {
			return 1
		}
		
		atime := utils.GetUnixTime(athleteA.FinishLineTime)
		btime := utils.GetUnixTime(athleteB.FinishLineTime)

		return utils.MaxTime(atime, btime)	
		
	}

	if !utils.IsEmptyString(athleteB.FinishLineTime) {
		return -1
	}

	atime := utils.GetUnixTime(athleteA.FinishCorridorTime)
	btime := utils.GetUnixTime(athleteB.FinishCorridorTime)

	return utils.MaxTime(atime, btime)	
}

// the order is ascending
var finishCorridorRecords = treeset.NewWith(byTimePoint)
// the order is ascending
var finishLineRecords =  treeset.NewWith(byTimePoint)



// GetRecords return sorted records of the athletes
func GetRecords(id int, recordType pb.TimingSystemRequest_FinishTypes, timepoint string) Athletes {

	sportsmanInfo := getAthleteInfo(id)

	switch recordType {
	case pb.TimingSystemRequest_FINISH_CORRIDOR:
		finishCorridorRecords.Add(athlete{
			ID: id,
			FinishCorridorTime: timepoint,
			FinishLineTime: "",
			FullName: sportsmanInfo.FullName,
			StartNumber: sportsmanInfo.StartNumber,
		})

	case pb.TimingSystemRequest_FINISH_LINE:
		finishLineRecords.Add(athlete{
			ID: id,
			FinishCorridorTime: "",
			FinishLineTime: timepoint,
			FullName: sportsmanInfo.FullName,
			StartNumber: sportsmanInfo.StartNumber,
		})
	
	default:
		log.Fatalf("The finish type is wrong!")
	}

	athleteRanking := mergeRecords()

	if len(athleteRanking) == 0 {
		log.Fatalf("merge finishline and finishcorridor failed")
	}
	return athleteRanking
}


func getAthleteInfo(athleteID int) athleteInfo{
	db, err := sql.Open("sqlite3", "../athletes.db")
	cerror.CheckErr(err)
	defer db.Close()

	selectSQL := "SELECT * FROM athletes WHERE Id = ?"
	stmt, err := db.Prepare(selectSQL)
	cerror.CheckErr(err)
	defer stmt.Close()

	var id int
	var fullName string
	var startNumber int
	err = stmt.QueryRow(athleteID).Scan(&id, &fullName, &startNumber)
	cerror.CheckErr(err)

	return athleteInfo {
		ID: id,
		StartNumber: startNumber,
		FullName: fullName,
	}

}

func mergeRecords() []athlete {
	athleteRanking := make([]athlete, 0)

	// make a copy of FinishCorridorRecords
	fcRecords := treeset.NewWith(byTimePoint)
	it := finishCorridorRecords.Iterator()
	for it.Next() {
		fcRecords.Add(it.Value())
	}

	// merge the records of finishline into the athlete ranking
	it = finishLineRecords.Iterator()
	for it.Next() {

		_, v := it.Index(), it.Value()
		sportsman := v.(athlete)

		_, theMan := finishCorridorRecords.Find(func (index int, elm interface{}) bool {
			r := elm.(athlete)
			return r.ID == sportsman.ID
		})

		if theMan == nil {
			log.Fatalf("Cannot find the ID:%v in the finishCorridorRecords", sportsman.ID)
		}
		info := theMan.(athlete)

		sportsman.FinishCorridorTime = info.FinishCorridorTime
		sportsman.FullName = info.FullName
		sportsman.StartNumber = info.StartNumber

		fcRecords.Remove(theMan)

		athleteRanking = append(athleteRanking, sportsman)
	}

	// merge the records of the finishcorridor into the athlete ranking
	if fcRecords.Size() != 0 {
		it = fcRecords.Iterator()
		for it.Next() {
			_, v := it.Index(), it.Value()
			sportsman := v.(athlete)

			athleteRanking = append(athleteRanking, sportsman)
		}
	}

	return athleteRanking
}


