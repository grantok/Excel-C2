package C2

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
)

// HTTPClient interface
type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

type Client struct {
	// BaseURL      *url.URL
	BaseURL      string
	UserAgent    string
	TenantId     string
	ClientId     string
	ClientSecret string
	DriveId      string
	SheetId      string
	SheetName    string
	TokenId      string
	HttpClient   HTTPClient
	APIKey       string
}

func (c *Client) newRequest(method, path string, body interface{}) (*http.Request, error) {
	if c.BaseURL == "" {
		return nil, errors.New("BaseURL is undefined")
	}

	// rel := &url.URL{Path: path}
	// u := c.BaseURL.ResolveReference(rel)
	u, _ := url.JoinPath(c.BaseURL, path)

	// var buf io.ReadWriter
	// if body != nil {
	// 	buf = new(bytes.Buffer)
	// 	err := json.NewEncoder(buf).Encode(body)
	// 	if err != nil {
	// 		fmt.Println("Error encoding body")
	// 		return nil, err
	// 	}
	// }
	// body_string := body.(string)
	// buf := bytes.NewBufferString(body_string)
	buf := body.(*bytes.Buffer)
	req, err := http.NewRequest(method, u, buf)
	// req, err := http.NewRequest(method, u.String(), buf)
	if err != nil {
		return nil, err
	}
	// Default request is json
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("User-Agent", c.UserAgent)
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.APIKey))
	return req, nil
}

func (c *Client) do(req *http.Request,
	v interface{}) (*http.Response, error) {
	resp, err := c.HttpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	err = json.NewDecoder(resp.Body).Decode(v)

	return resp, err
}
