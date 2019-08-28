package data

import (
	"os"
	"database/sql"
	
	// import go-sqlite3 driver
	_ "github.com/mattn/go-sqlite3"
	"timingsystem/timingserver/cerror"
)

var athletes = make([]athleteInfo, 0)
// InitialAthletesData initials athleles data
func initialAthletesData() {
	
	athletes = append(athletes, athleteInfo{ ID: 1, StartNumber: 100, FullName: "AthleteA" })
	athletes = append(athletes, athleteInfo{ ID: 2, StartNumber: 101, FullName: "AthleteB" })
	athletes = append(athletes, athleteInfo{ ID: 3, StartNumber: 102, FullName: "AthleteC" })
	athletes = append(athletes, athleteInfo{ ID: 4, StartNumber: 103, FullName: "AthleteD" })
	athletes = append(athletes, athleteInfo{ ID: 5, StartNumber: 104, FullName: "AthleteE" })
	athletes = append(athletes, athleteInfo{ ID: 6, StartNumber: 105, FullName: "AthleteF" })
	athletes = append(athletes, athleteInfo{ ID: 7, StartNumber: 106, FullName: "AthleteG" })
	athletes = append(athletes, athleteInfo{ ID: 8, StartNumber: 107, FullName: "AthleteH" })
	athletes = append(athletes, athleteInfo{ ID: 9, StartNumber: 108, FullName: "AthleteI" })
	athletes = append(athletes, athleteInfo{ ID: 10, StartNumber: 109, FullName: "AthleteJ" })
}

// InitialDatabase inserts the athletes data into database
func InitialDatabase() bool {
	os.Remove("../athletes.db")

	initialAthletesData()

	db, err := sql.Open("sqlite3", "../athletes.db")
	cerror.CheckErr(err)
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
	cerror.CheckErr(err)

	insertSQL := "INSERT INTO athletes(Id, FullName, StartNumber) values(?, ?, ?)"
	stmt, err := db.Prepare(insertSQL)
	cerror.CheckErr(err)
	defer stmt.Close()

	for _, value := range athletes {
		_, err = stmt.Exec(value.ID, value.FullName, value.StartNumber)
		cerror.CheckErr(err)
	}

	return true
}