package spark

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/url"
	"time"
)

const MessagesUrl = "https://api.ciscospark.com/v1/messages"

type Message struct {
	Id          string    `json:"id,omitempty"`
	RoomId      string    `json:"roomId,omitempty"`
	RoomType    string    `json:"roomType,omitempty"`
	Text        string    `json:"text,omitempty"`
	PersonId    string    `json:"personId,omitempty"`
	PersonEmail string    `json:"personEmail,omitempty"`
	Markdown    string    `json:"markdown,omitempty"`
	Html        string    `json:"html,omitempty"`
	created     time.Time `json:"created,omitempty"`
}

type Messages struct {
	Items []Message
}

// url.Values can be found here:
// https://developer.ciscospark.com/endpoint-messages-get.html
// example:
// uv := url.Values{}
// uv.Add("type", "group")
// s.ListRooms(&uv)
func (s *Spark) ListMessages(uv *url.Values) ([]Message, error) {
	var m Messages
	if uv == nil {
		return m.Items, errors.New("Please include query params")
	}
	if (*uv).Get("roomId") == "" {
		return m.Items, errors.New("Please include a roomId")
	}

	bytes, err := s.GetRequest(MessagesUrl, uv)
	if err != nil {
		return m.Items, err
	}
	//fmt.Println(string(bytes))
	err = json.Unmarshal(bytes, &m)
	return m.Items, err
}

func (s *Spark) CreateMessage(m Message) (Message, error) {
	var rm Message
	if m.RoomId == "" {
		return rm, errors.New("Please include a roomId")
	}

	if m.RoomType == "" {
		m.RoomType = "group"
	}

	b := new(bytes.Buffer)
	json.NewEncoder(b).Encode(m)
	bytes, err := s.PostRequest(MessagesUrl, b)

	if err != nil {
		return rm, err
	}
	err = json.Unmarshal(bytes, &rm)
	return rm, err
}
