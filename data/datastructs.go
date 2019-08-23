package data

type athlete struct {
	ID int
	FullName string
	StartNumber int
	FinishCorridorTime string
	FinishLineTime string
}

type athleteInfo struct {
	ID int
	StartNumber int
	FullName string
}

// Athletes is a set of athlete
type Athletes []athlete