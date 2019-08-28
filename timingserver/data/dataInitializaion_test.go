package data

import (
	"testing"
)

func TestInitialDatabase(t *testing.T) {
	res := InitialDatabase()
	
	if res {
		t.Log("Initial database successful!")
	} else {
		t.Error("Initial database failed!")
	}
}