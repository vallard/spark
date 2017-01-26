package spark

import (
	"encoding/json"
	"errors"
	"fmt"
	"time"
)

const PeopleUrl = "https://api.ciscospark.com/v1/people"

type Person struct {
	Id           string    `json:"id"`
	Emails       []string  `json:"emails"`
	DisplayName  string    `json:"displayName"`
	FirstName    string    `json:"firstName"`
	LastName     string    `json:"lastName"`
	Avatar       string    `json:"avatar"`
	OrgId        string    `json:"orgId"`
	Roles        []string  `json:"roles"`
	Licenses     []string  `json:"licenses"`
	Created      time.Time `json:"created"`
	Timezone     string    `json:"timezone"`
	LastActivity time.Time `json:"lastActivity"`
	Status       string    `json:"status"`
}

// https://developer.ciscospark.com/resource-webhooks.html

type People struct {
	Items []Person
}

func (s *Spark) GetPerson(personId string) (Person, error) {
	var person Person
	if personId == "" {
		return person, errors.New("No person ID specified")
	}
	bytes, err := s.GetRequest(fmt.Sprintf("%s/%s", PeopleUrl, personId), nil)
	if err != nil {
		return person, err
	}
	err = json.Unmarshal(bytes, &person)
	return person, err
}

func (s *Spark) GetMyself() (Person, error) {
	var me Person
	bytes, err := s.GetRequest(fmt.Sprintf("%s/me", PeopleUrl), nil)
	if err != nil {
		return me, err
	}
	err = json.Unmarshal(bytes, &me)
	return me, err
}
