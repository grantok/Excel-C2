package C2

import (
	"log"
)

type WorkSheet struct {
	Id         string `json:"id"`
	Position   int    `json:"position"`
	Name       string `json:"name"`
	Visibility string `json:"visibility"`
}

func (c *Client) AddSheet(name string) (*WorkSheet, error) {

	body := `{"name":"` + name + `"}`
	// body := []byte(`{"name": "` + name + `"}`)
	req, err := c.newRequest("POST", "/add", body)
	if err != nil {
		log.Fatal(err)
	}

	ws := new(WorkSheet)
	_, err = c.do(req, &ws)
	if err != nil {
		return nil, err
	}
	return ws, err
}
