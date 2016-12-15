package spark

import (
	"os"
	"testing"
)

var s *Spark
var allRooms []SparkRoom

// initialize spark client
func GetSpark(t *testing.T) *Spark {
	token := os.Getenv("SPARK_TOKEN")
	if token == "" {
		t.Fatal("Please define the SPARK_TOKEN environment variable before running tests")
	}
	return New(token)
}

// test if there are rooms
func TestListRooms(t *testing.T) {
	s = GetSpark(t)
	rooms, err := s.ListRooms(nil)
	if err != nil {
		t.Error(err)
	}
	if len(rooms) < 1 {
		t.Fatalf("expected at least one room to be listed.  Instead got %d\n", len(rooms))
	}
	allRooms = rooms

}

// get the first room
func TestGetRoom(t *testing.T) {
	s = GetSpark(t)
	r := allRooms[0].Id
	room, err := s.GetRoom(r)
	if err != nil {
		t.Error(err)
	}
	if room.Id != r {
		t.Error("Expected room ID to be the same")
	}
}
