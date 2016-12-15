package spark

import (
	"fmt"
	"net/url"
	"testing"
)

func TestBadMessageRequest(t *testing.T) {
	s = getSpark(t)
	_, err := s.ListMessages(nil)
	if err == nil {
		t.Error("Should get an error when making a request for messages without specifying the room")
	}
}

// test if there are rooms
func TestListMessages(t *testing.T) {
	s = getSpark(t)
	if len(allRooms) < 1 {
		getAllRooms(t)
	}

	uv := url.Values{}
	uv.Add("roomId", allRooms[0].Id)
	m, err := s.ListMessages(&uv)

	if err != nil {
		t.Error(err)
	}

	if len(m) < 1 {
		t.Error("Expected there to be at least one message in the room")
	}
}

func TestCreateMessage(t *testing.T) {
	s = getSpark(t)
	if len(allRooms) < 1 {
		getAllRooms(t)
	}

	room := allRooms[0]
	m := Message{
		RoomId: room.Id,
		Text:   "sorry, just a quick test",
	}
	rm, err := s.CreateMessage(m)
	if err != nil {
		t.Error(err)
	}
	if m.Text != rm.Text {
		t.Errorf("titles of return message %s should be the same as original: %s", m.Text, rm.Text)
	}
	fmt.Println(rm)
}
