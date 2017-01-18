package spark

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/url"
)

const WebhookUrl = "https://api.ciscospark.com/v1/webhooks"

type Webhook struct {
	Id        string                 `json:"id"`
	Name      string                 `json:"name"`
	TargetUrl string                 `json:"targetUrl"`
	Resource  string                 `json:"resource"`
	Event     string                 `json:"event"`
	Filter    string                 `json:"filter,omitempty"`
	Secret    string                 `json:"secret,omitempty"`
	OrgId     string                 `json:"orgId,omitempty"`
	CreatedBy string                 `json:"createdBy,omitempty"`
	AppId     string                 `json:"appId,omitempty"`
	OwnedBy   string                 `json:"ownedBy,omitempty"`
	Status    string                 `json:"active,omitempty"`
	ActorId   string                 `json:"actorId,omitempty"`
	Data      map[string]interface{} `json:"data,omitempty"`
}

// https://developer.ciscospark.com/resource-webhooks.html

type Webhooks struct {
	Items []Webhook
}

func (s *Spark) ListWebhooks(uv *url.Values) ([]Webhook, error) {
	var w Webhooks
	// parameter: max is the only accepted right now.
	// https://developer.ciscospark.com/endpoint-webhooks-get.html

	bytes, err := s.GetRequest(WebhookUrl, uv)
	if err != nil {
		return w.Items, err
	}

	err = json.Unmarshal(bytes, &w)
	return w.Items, err
}

func (s *Spark) CreateWebhook(w Webhook) (Webhook, error) {
	var rwh Webhook
	if w.Name == "" {
		return rwh, errors.New("You must specify a name for the webhook")
	}
	if w.TargetUrl == "" {
		return rwh, errors.New("You must specify a target URL for the webhook")
	}
	// see: https://developer.ciscospark.com/webhooks-explained.html
	// resource should be the plural form of a spark api:
	// * messages
	// * memberships
	// * rooms
	// * teams
	if w.Resource == "" {
		return rwh, errors.New("You must specify a Resource for the webhook")
	}

	// created, updated, deleted, or all
	if w.Event == "" {
		return rwh, errors.New("You should specify an event for the webhook resource")
	}

	/* generally, you'll probably want:
	{"resource" : "messages", "event" : "created" }
	*/
	b := new(bytes.Buffer)
	json.NewEncoder(b).Encode(w)
	bytes, err := s.PostRequest(WebhookUrl, b)

	if err != nil {
		return rwh, err
	}
	err = json.Unmarshal(bytes, &rwh)
	return rwh, err
}

func (s *Spark) DeleteWebhook(w Webhook) error {
	if w.Id == "" {
		return errors.New("Must specify the Webhook ID to delete")
	}

	url := WebhookUrl + "/" + w.Id
	_, err := s.DeleteRequest(url)
	return err
}
