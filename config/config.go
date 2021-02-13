package config

import (
	secretmanager "cloud.google.com/go/secretmanager/apiv1"
	"cloudhired.com/api/clogger"
	"context"
	secretmanagerpb "google.golang.org/genproto/googleapis/cloud/secretmanager/v1"
)

const (
	CONNSTRING = "projects/782780515351/secrets/atlas-conn-string/versions/latest"
)

var ConnectionString string

func init() {
	ctx := context.Background()
	client, err := secretmanager.NewClient(ctx)
	if err != nil {
		clogger.Error(err)
	}

	// get atlas connection string
	req := &secretmanagerpb.AccessSecretVersionRequest{Name: CONNSTRING}
	result, err := client.AccessSecretVersion(ctx, req)
	if err != nil {
		clogger.Error(err)
	}
	ConnectionString = string(result.Payload.Data)
}
