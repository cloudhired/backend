package clogger

import (
	"cloud.google.com/go/logging"
	"context"
	"fmt"
	"log"
)

const (
	projectID = "782780515351"
	logName   = "cloudhired-backend-log"
)

//var client *logging.Client
var logger *logging.Logger

func init() {
	ctx := context.Background()
	client, err := logging.NewClient(ctx, projectID)
	if err != nil {
		log.Fatalf("Failed to create logging client: %v", err)
	}

	logger = client.Logger(logName)
	//.StandardLogger(logging.Error)
	//logger.Println("hello world..")
	//fmt.Println("can you see?")
	//defer client.Close()
}

func LogError(err error) {
	fmt.Println("before call")
	errorLogger := logger.StandardLogger(logging.Error)
	fmt.Println("after call")
	errorLogger.Println("error msg:", err)
	//defer client.Close()
}
