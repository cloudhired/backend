package config

import (
	secretmanager "cloud.google.com/go/secretmanager/apiv1"
	"cloudhired.com/api/clogger"
	"context"
	secretmanagerpb "google.golang.org/genproto/googleapis/cloud/secretmanager/v1"
)

const (
	CONNSTRING = "projects/782780515351/secrets/my-secret/versions/latest"
)

var ConnectionString string

func init() {
	ctx := context.Background()
	client, err := secretmanager.NewClient(ctx)
	if err != nil {
		clogger.Error(err)
	}

	clogger.Default("default message@@@@@@@@@@@@")
	clogger.Alert("alert message")
	clogger.Critical("very critical")
	clogger.Info("just info")
	clogger.Emergency("This is emergency. DO NOT ignore!!")
	clogger.Warning("warning!!")
	clogger.Notice("just a notice")
	clogger.Debug("debugging...")

	req := &secretmanagerpb.AccessSecretVersionRequest{
		Name: CONNSTRING,
	}
	result, err := client.AccessSecretVersion(ctx, req)
	if err != nil {
		clogger.Error(err)
	}
	ConnectionString = string(result.Payload.Data)
}
