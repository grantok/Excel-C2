package C2

import (
	"Excel-C2/internal/authentication"
	"Excel-C2/internal/configuration"
	"context"
	"fmt"
)

func Run() {

	// perform authentication
	_, graph_client := authentication.Authenticate(configuration.GetOptionsTenantId(),
		configuration.GetOptionsClientId(),
		configuration.GetOptionsClientSecret())

	// sheet configuration
	spreadSheet := &configuration.SpreadSheet{}
	spreadSheet.SpreadSheetId = configuration.GetOptionsSheetId()
	spreadSheet.DriveId = configuration.GetOptionsDriveId()

	// create new sheet ?

	// ticker

	// TEST
	drive_item, err := graph_client.Drives().
		ByDriveId(spreadSheet.DriveId).
		Items().
		ByDriveItemId(spreadSheet.SpreadSheetId).
		Get(context.Background(), nil)
	if err != nil {
		fmt.Printf("Error downloading workbook: %v\n", err)
	} else {
		fmt.Println("dowload succeeded")
		// "@microsoft.graph.downloadUrl" from json will contain a URL to grab the file
		fmt.Println(drive_item.GetBackingStore().Enumerate())
		fmt.Printf("drive_item: %v\n", drive_item)
		ad := drive_item.GetAdditionalData()["@microsoft.graph.downloadUrl"]
		fmt.Println(*ad.(*string))
	}
}
