package config

import (
	"context"
	"fmt"
	"log"

	secretmanager "cloud.google.com/go/secretmanager/apiv1"
	secretmanagerpb "google.golang.org/genproto/googleapis/cloud/secretmanager/v1"
)

const (
	ConnectionString = "mongodb+srv://ch-user:FhDne1WoX3qI2wIm@cloudhired.c58f7.gcp.mongodb.net/cloudhired?retryWrites=true&w=majority"
)

func init() {
	ctx := context.Background()
	client, err := secretmanager.NewClient(ctx)
	fmt.Println("here")
	if err != nil {
		log.Fatal("client: ", err)
	}

	req := &secretmanagerpb.AccessSecretVersionRequest{
		Name: "projects/782780515351/secrets/my-secret/versions/latest",
	}

	result, err := client.AccessSecretVersion(ctx, req)
	if err != nil {
		log.Fatal("the Error: ", err)
	}

	log.Fatal("result: ", string(result.Payload.Data))
}
