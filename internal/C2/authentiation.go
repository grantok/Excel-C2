package C2

import (
	"Excel-C2/internal/utils"
	"context"
	"fmt"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	azidentity "github.com/Azure/azure-sdk-for-go/sdk/azidentity"
)

// func (c *Client) Authenticate() (context.Context, *msgraphsdk.GraphServiceClient) {
func (c *Client) Authenticate() (context.Context, string) {
	ctx := context.Background()

	//retrieve the credential
	cred, err := azidentity.NewClientSecretCredential(c.TenantId, c.ClientId, c.ClientSecret, nil)

	if err != nil {
		utils.LogFatalDebug("[-] Authentication failed")
	} else {
		utils.LogDebug("Auth success")
	}

	tkn, err := cred.GetToken(ctx, policy.TokenRequestOptions{
		Scopes: []string{
			"https://graph.microsoft.com/.default",
		},
	})
	if err != nil {
		utils.LogFatalDebug("[-] Get token failed")
	} else {
		utils.LogDebug("token success")
	}
	c.APIKey = tkn.Token
	fmt.Println(tkn.Token)

	return ctx, tkn.Token
}
