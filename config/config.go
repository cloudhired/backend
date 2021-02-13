package config

import (
	secretmanager "cloud.google.com/go/secretmanager/apiv1"
	"cloudhired.com/api/clogger"
	"context"
	"fmt"
	secretmanagerpb "google.golang.org/genproto/googleapis/cloud/secretmanager/v1"
	"log"
)

var ConnectionString string

func init() {
	ctx := context.Background()
	client, err := secretmanager.NewClient(ctx)
	if err != nil {
		clogger.LogError(err)
	}

	req := &secretmanagerpb.AccessSecretVersionRequest{
		Name: "projects/782780515351/secrets/my-secret/versions/2",
	}

	result, err := client.AccessSecretVersion(ctx, req)
	if err != nil {
		log.Fatal("the Error: ", err)
	}

	ConnectionString = string(result.Payload.Data)

	fmt.Println("secret result: ", string(result.Payload.Data))
}
