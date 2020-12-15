package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	Username		primitive.ObjectID	`bson:"_id" json:"id"`
	FullName		string				`bson:"fullname" json:"fullname"`
	FirstName		string				`bson:"fname" json:"fname"`		
}

