package C2

import (
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
		c.LogFatalDebug("[-] Authentication failed")
	} else {
		c.LogDebug("Auth success")
	}

	tkn, err := cred.GetToken(ctx, policy.TokenRequestOptions{
		Scopes: []string{
			"https://graph.microsoft.com/.default",
		},
	})
	if err != nil {
		c.LogFatalDebug("[-] Get token failed")
	} else {
		c.LogDebug("token success")
	}
	c.APIKey = tkn.Token
	fmt.Println(tkn.Token)

	return ctx, tkn.Token
}
