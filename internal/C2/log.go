package C2

import "log"

func (c *Client) LogDebug(message string) {
	if c.Debug {
		log.Println(message)
	}
}

func (c *Client) LogFatalDebug(message string) {
	if c.Debug {
		log.Fatal(message)
	}
}

func (c *Client) LogFatalDebugError(message string, err error) {
	if c.Debug {
		log.Fatal(message, err)
	}
}
