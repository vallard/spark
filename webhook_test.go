package spark

import (
	"fmt"
	"net/url"
	"testing"
)

func TestListWebhooks(t *testing.T) {
	s = getSpark(t)
	if len(allRooms) < 1 {
		getAllRooms(t)
	}
	wh, err := s.ListWebhooks(&url.Values{})
	if err != nil {
		t.Error(err)
	}
	fmt.Println("All webhooks: ", wh)
}

func TestCreateDeleteWebhook(t *testing.T) {
	s = getSpark(t)
	w := Webhook{
		Name:      "Test",
		TargetUrl: "https://bots.ciscopipeline.io/vallard/rickjames",
		Resource:  "messages",
		Event:     "created",
	}
	fmt.Println("Creating this new webhook:", w)
	rw, err := s.CreateWebhook(w)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println("webhook created: ", rw)

	err = s.DeleteWebhook(rw)
	if err != nil {
		t.Error(err)
	}
}
