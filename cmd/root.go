package cmd

import (
	"os"

	"Excel-C2/internal/C2"
	"Excel-C2/internal/configuration"

	"github.com/spf13/cobra"
)

var (
	tenantId     string
	clientId     string
	clientSecret string
	driveId      string
	sheetId      string
	debug        bool
)

var rootCmd = &cobra.Command{
	Use:   "excel-c2",
	Short: "excel-c2",
	Long:  `excel-c2`,

	Run: func(cmd *cobra.Command, args []string) {
		configuration.SetOptions(tenantId, clientId, clientSecret, driveId, sheetId, debug)
		C2.Run()
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
