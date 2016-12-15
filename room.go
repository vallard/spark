package spark

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/url"
	"time"
)

const (
	RoomsUrl = "https://api.ciscospark.com/v1/rooms"
)

type SparkRoom struct {
	Id           string
	Title        string
	Type         string
	IsLocked     bool
	LastActivity time.Time
	CreatorId    string
	Created      time.Time
}

type SparkRooms struct {
	Items []SparkRoom
}

// List the rooms.  URL values that can be used are:
// teamId string
// max Integer
// type string that is either 'group' or 'string'
// uv := url.Values{}
// uv.Add("type", "group")
// s.ListRooms(&uv)
func (s *Spark) ListRooms(uv *url.Values) ([]SparkRoom, error) {
	bytes, err := s.GetRequest(RoomsUrl, uv)
	var rooms SparkRooms
	if err != nil {
		return rooms.Items, err
	}
	err = json.Unmarshal(bytes, &rooms)
	return rooms.Items, nil
}

// get one room
func (s *Spark) GetRoom(roomId string) (SparkRoom, error) {
	var room SparkRoom
	if roomId == "" {
		return room, errors.New("No Room ID was specified")
	}
	bytes, err := s.GetRequest(fmt.Sprintf("%s/%s", RoomsUrl, roomId), nil)
	if err != nil {
		return room, err
	}
	err = json.Unmarshal(bytes, &room)
	return room, err
}
