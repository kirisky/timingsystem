package data

// athlete information
type athlete struct {
	ID int
	FullName string
	StartNumber int
	FinishCorridorTime string
	FinishLineTime string
}

// Athletes is a set of athlete
type Athletes []athlete

type athleteInfo struct {
	ID int
	StartNumber int
	FullName string
}
