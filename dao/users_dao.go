package dao

import (
	"cloudhired.com/api/helper"
	"cloudhired.com/api/models"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"strings"
)

const (
	CONNECTIONSTRING = "mongodb+srv://ch-user:FhDne1WoX3qI2wIm@cloudhired.c58f7.gcp.mongodb.net/cloudhired?retryWrites=true&w=majority"
	DBNAME           = "cloudhired"
	COLLECTION       = "users"
)

var db *mongo.Database
var collection *mongo.Collection

func init() {
	clientOptions := options.Client().ApplyURI(CONNECTIONSTRING)
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	// Collection types can be used to access the database
	db = client.Database(DBNAME)
	collection = db.Collection(COLLECTION)
}

func GetUsernameByUid(uid string, name string, email string, emailVerifiled bool) string {
	filter := bson.D{
		{"uid", uid},
		{"email", email},
	}
	projection := bson.D{
		{"username", 1},
	}
	var u bson.M
	err := collection.FindOne(context.Background(), filter, options.FindOne().SetProjection(projection)).Decode(&u)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			// if no doc found, create new user
			usernameList := strings.Split(strings.TrimSpace(strings.ToLower(name)), " ")
			usernameList = append(usernameList, helper.RandomString(8))
			username := strings.Join(usernameList, "-")
			insert := bson.D{
				{"uid", uid},
				{"username", username},
				{"email", email},
				{"fullname", name},
				{"email_verified", emailVerifiled},
			}
			_, insErr := collection.InsertOne(context.Background(), insert)
			if insErr != nil {
				log.Fatal(insErr)
			}
			return username
		}
		fmt.Println("error occured")
		log.Fatal(err)
	}
	return u["username"].(string)
}

func AllUsers() []models.User {
	findOptions := options.Find()
	findOptions.SetLimit(10)
	cur, err := collection.Find(context.Background(), bson.D{{}})
	if err != nil {
		log.Fatal(err)
	}
	var elements []models.User
	var e models.User
	// Get the next result from the cursor
	for cur.Next(context.Background()) {
		err := cur.Decode(&e)
		if err != nil {
			log.Fatal(err)
		}
		elements = append(elements, e)
	}
	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}
	cur.Close(context.Background())
	return elements
}

func FindOneUser(u string) models.User {
	var user models.User
	filter := bson.D{{"username", u}}
	err := collection.FindOne(context.Background(), filter).Decode(&user)
	if err != nil {
		log.Fatal(err)
	}
	return user
}

func UpdateOneUser(u string, doc interface{}) bool {
	opts := options.Update().SetUpsert(false)
	filter := bson.D{{"username", u}}
	update := bson.D{{"$set", doc}}
	_, err := collection.UpdateOne(context.TODO(), filter, update, opts)
	if err != nil {
		log.Fatal(err)
		return false
	}
	return true
}

func toDoc(v interface{}) (doc *bson.D, err error) {
	data, err := bson.Marshal(v)
	if err != nil {
		return
	}
	err = bson.Unmarshal(data, &doc)
	return
}

func TestDb() {
	collection := db.Collection(COLLECTION)
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
