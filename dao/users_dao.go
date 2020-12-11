package dao

import (
	"context"
	"fmt"
	"log"
	// "cloudhired.com/api/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	CONNECTIONSTRING = "mongodb+srv://ch-user:FhDne1WoX3qI2wIm@cloudhired.c58f7.gcp.mongodb.net/cloudhired?retryWrites=true&w=majority"
	DBNAME = ""
	COLLECTION =  "users"
)

var db *mongo.Database

// Connect establish a connection to database
func init() {
	clientOptions := options.Client().ApplyURI(CONNECTIONSTRING)
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	// Collection types can be used to access the database
	db = client.Database("cloudhired")
}

func TestDb() {
	collection := db.Collection("users")
	findOptions := options.Find()
	findOptions.SetLimit(10)
	cur, err := collection.Find(context.TODO(), bson.D{{}})
	if err != nil {
		log.Fatal(err)
	}

	var results []bson.M
	if err = cur.All(context.TODO(), &results); err != nil {
		log.Fatal(err)
	}
	for _, result := range results {
		fmt.Println(result)
	}

	for cur.Next(context.TODO()) {
		var elem bson.M
		fmt.Println(elem)
	}
}

