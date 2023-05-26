package C2

import "log"

func (c *Client) LogDebug(message string) {
	// if configuration.GetOptionsDebug() {
	log.Println(message)
	// }
}

func (c *Client) LogFatalDebug(message string) {
	// if configuration.GetOptionsDebug() {
	log.Fatal(message)
	// }
}

func (c *Client) LogFatalDebugError(message string, err error) {
	// if configuration.GetOptionsDebug() {
	log.Fatal(message, err)
	// }
}
