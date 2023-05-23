package authentication

import (
	"Excel-C2/internal/utils"
	"context"
	"fmt"

	azidentity "github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	msgraphsdk "github.com/microsoftgraph/msgraph-sdk-go"
)

func Authenticate(tenant string, client string, secret string) (context.Context, *msgraphsdk.GraphServiceClient) {
	ctx := context.Background()

	//retrieve the credential
	cred, err := azidentity.NewClientSecretCredential(tenant, client, secret, nil)

	if err != nil {
		utils.LogFatalDebug("[-] Authentication failed")
	} else {
		fmt.Println("Auth success")
	}

	//retrieve the client
	graphclient, err := msgraphsdk.NewGraphServiceClientWithCredentials(cred, []string{".default"})
	if err != nil {
		utils.LogFatalDebug("[-] Client failed")
	} else {
		fmt.Println("Client success")
	}

	return ctx, graphclient
}
