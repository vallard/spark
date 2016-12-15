package spark

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

const (
	MessagesUrl = "https://api.ciscospark.com/v1/messages"
	PeopleUrl   = "https://api.ciscospark.com/v1/people"
)

var httpClient *http.Client

func init() {
	httpClient = &http.Client{}
}

func (s *Spark) request(req *http.Request) ([]byte, error) {
	// set headers for all requests
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", s.token))
	req.Header.Set("Content-Type", "application/json; charset=utf-8")

	res, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	return ioutil.ReadAll(res.Body)
}

func (s *Spark) GetRequest(url string, uv *url.Values) ([]byte, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	if uv != nil {
		req.URL.RawQuery = (*uv).Encode()
	}
	return s.request(req)
}

func (s *Spark) PostRequest(url string, uv *url.Values, body *bytes.Buffer) ([]byte, error) {
	req, err := http.NewRequest("POST", url, body)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	if err != nil {
		return nil, err
	}
	if uv != nil {
		req.URL.RawQuery = (*uv).Encode()
	}
	return s.request(req)
}
