package spark

import (
	"os"
	"testing"
)

var s *Spark
var allRooms []SparkRoom

// initialize spark client
func getSpark(t *testing.T) *Spark {
	token := os.Getenv("SPARK_TOKEN")
	if token == "" {
		t.Fatal("Please define the SPARK_TOKEN environment variable before running tests")
	}
	return New(token)
}

func getAllRooms(t *testing.T) {
	rooms, err := s.ListRooms(nil)
	if err != nil {
		t.Error(err)
	}
	if len(rooms) < 1 {
		t.Fatalf("expected at least one room to be listed.  Instead got %d\n", len(rooms))
	}
	allRooms = rooms

}

// test if there are rooms
func TestListRooms(t *testing.T) {
	s = getSpark(t)
	if len(allRooms) < 1 {
		getAllRooms(t)
	}
}

// get the first room
func TestGetRoom(t *testing.T) {
	s = getSpark(t)
	r := allRooms[0].Id
	room, err := s.GetRoom(r)
	if err != nil {
		t.Error(err)
	}
	if room.Id != r {
		t.Error("Expected room ID to be the same")
	}
}

func TestGetRoomName(t *testing.T) {
	s = getSpark(t)
	title := allRooms[0].Title
	room, err := s.GetRoomWithName(title)
	if err != nil {
		t.Error(err)
	}
	if room.Title != title {
		t.Error("Expected the room Name to be the same as what we just got")
	}

	// expecting this room to be nil
	room, err = s.GetRoomWithName("Fizzle Swamp Duck Chocolate")
	if err == nil {
		t.Error("Rooms that are not found should return an error")
	}
	if room.Title != "" {
		t.Error("Rooms that are not found should be empty")
	}
}
