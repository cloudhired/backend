/*
The purpose of this package is to have a wrapper for GCP logging package.
When needed in other package, we can just import this one without having to create new client every time.

Written by Morgan Gao
2/12/2021
*/

package clogger

import (
	"cloud.google.com/go/logging"
	"context"
	"log"
)

const (
	projectID = "782780515351"
	logName   = "cloudhired-backend-log"
)

var logger *logging.Logger

func init() {
	ctx := context.Background()

	// initialize gcp logging client
	client, err := logging.NewClient(ctx, projectID)
	if err != nil {
		log.Fatalf("Failed to create logging client: %v", err)
	}
	logger = client.Logger(logName)
}

func Default(msg string) {
	defaultLogger := logger.StandardLogger(logging.Default)
	defaultLogger.Println(msg)
}

func Alert(alert string) {
	alertLogger := logger.StandardLogger(logging.Alert)
	alertLogger.Println(alert)
}

func Error(err error) {
	errorLogger := logger.StandardLogger(logging.Error)
	errorLogger.Println(err)
}

func Info(info string) {
	infoLogger := logger.StandardLogger(logging.Info)
	infoLogger.Println(info)
}

func Critical(info string) {
	infoLogger := logger.StandardLogger(logging.Critical)
	infoLogger.Println(info)
}

func Emergency(msg string) {
	emergLogger := logger.StandardLogger(logging.Emergency)
	emergLogger.Println(msg)
}

func Warning(msg string) {
	warnLogger := logger.StandardLogger(logging.Warning)
	warnLogger.Println(msg)
}

func Notice(msg string) {
	noticeLogger := logger.StandardLogger(logging.Notice)
	noticeLogger.Println(msg)
}

func Debug(msg string) {
	debugLogger := logger.StandardLogger(logging.Debug)
	debugLogger.Println(msg)
}
