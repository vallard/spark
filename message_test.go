package spark

import (
	"net/url"
	"testing"
)

var allMessages []Message

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
	allMessages, err := s.ListMessages(&uv)

	if err != nil {
		t.Error(err)
	}

	if len(allMessages) < 1 {
		t.Fatal("Expected there to be at least one message in the room")

	}

	// test Get messages
	firstMessage := allMessages[0]
	_, err = s.GetMessage(firstMessage.Id)
	if err != nil {
		t.Error(err)
	}
}

func TestCreateMessage(t *testing.T) {
	s = getSpark(t)
	if len(allRooms) < 1 {
		getAllRooms(t)
	}

	room := allRooms[1]
	m := Message{
		RoomId: room.Id,
		Text:   "I'll test as much as I want. I'm Rick James.  Get up!",
	}
	rm, err := s.CreateMessage(m)
	if err != nil {
		t.Error(err)
	}
	if m.Text != rm.Text {
		t.Errorf("titles of return message %s should be the same as original: %s", m.Text, rm.Text)
	}
}
