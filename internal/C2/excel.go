package C2

import (
	"bytes"
	"fmt"
	"log"
)

type WorkSheet struct {
	Id         string `json:"id"`
	Position   int    `json:"position"`
	Name       string `json:"name"`
	Visibility string `json:"visibility"`
}

type CellRange struct {
	Address       string `json:"address"`
	AddressLocal  string `json:"addressLocal-value"`
	CellCount     int    `json:"cellCount"`
	ColumnCount   int    `json:"columnCount"`
	ColumnHidden  bool   `json:"columnHidden"`
	ColumnIndex   int    `json:"columnIndex"`
	Formulas      []any  `json:"formulas"`
	FormulasLocal []any  `json:"formulasLocal"`
	FormulasR1C1  []any  `json:"formulasR1C1"`
	Hidden        bool   `json:"hidden"`
	NumberFormat  []any  `json:"numberFormat"`
	RowCount      int    `json:"rowCount"`
	RowHidden     bool   `json:"rowHidden"`
	RowIndex      int    `json:"rowIndex"`
	Text          []any  `json:"text"`
	ValueTypes    []any  `json:"valueTypes"`
	Values        []any  `json:"values"`
}

type Ok struct {
	Ok bool `json:"ok"`
}

func (c *Client) AddSheet(name string) (*WorkSheet, error) {

	// TODO - check if the sheet exists first??

	body := `{"name":"` + name + `"}`
	// req, err := c.newRequest("POST", "/add", body)
	// body := map[string]string{
	// 	"name": name,
	// }
	// json_body, _ := json.Marshal(body)
	req, err := c.newRequest("POST", "/add", bytes.NewBufferString(body))
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

func (c *Client) UpdateRange(cells string, values_string string) (*CellRange, error) {

	body := `{"values":` + values_string + `}`
	req, err := c.newRequest("PATCH", `/`+c.SheetName+`/range(address='`+cells+`')`, bytes.NewBufferString(body))
	if err != nil {
		log.Fatal(err)
	}

	cr := new(CellRange)
	_, err = c.do(req, &cr)
	if err != nil {
		return nil, err
	}
	fmt.Println("cr", cr.Address, cr.Values)
	return cr, err
}
