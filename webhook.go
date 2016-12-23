package spark

const WebhookUrl = "https://api.ciscospark.com/v1/webhooks"

type Webhook struct {
	Id        string                 `json:id`
	Name      string                 `json:name`
	Resource  string                 `json:resource`
	Event     string                 `json:event`
	Filter    string                 `json:filter`
	OrgId     string                 `json:orgId`
	CreatedBy string                 `json:createdBy`
	AppId     string                 `json:appId`
	OwnedBy   string                 `json:ownedBy`
	Status    string                 `json:active`
	ActorId   string                 `json:actorId`
	Data      map[string]interface{} `json:data`
}

// https://developer.ciscospark.com/resource-webhooks.html
