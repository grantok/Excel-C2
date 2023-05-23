package configuration

type SpreadSheet struct {
	DriveId       string
	SpreadSheetId string
	CommandSheet  Sheet
}

type Sheet struct {
	Name                     string //sheet name
	CommandsExecution        []Commands
	Ticker                   int
	RangeTickerConfiguration string
}

type Commands struct {
	RangeIn  string
	RangeOut string
	RangeId  string
	Input    string
	Output   string
}
