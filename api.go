package spark

import (
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

const (
	PeopleUrl = "https://api.ciscospark.com/v1/people"
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
	bs, err := ioutil.ReadAll(res.Body)

	if res.StatusCode != http.StatusOK {
		e := fmt.Sprintf("HTTP Status Code: %d\n%s", res.StatusCode, string(bs))
		return nil, errors.New(e)
	}

	return bs, err
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

func (s *Spark) PostRequest(url string, body *bytes.Buffer) ([]byte, error) {
	req, err := http.NewRequest("POST", url, body)
	if err != nil {
		return nil, err
	}
	return s.request(req)
}
