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
	c2.SheetName = newSheetName
	_, err := c2.AddSheet(newSheetName)
	if err != nil {
		fmt.Println("Add sheet error ", err)
	}

	_, err = c2.UpdateRange("A1", `[["Test"]]`)
	if err != nil {
		fmt.Println("Cell change error ", err)
	}

}
