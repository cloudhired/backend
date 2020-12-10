package dao

import (
	"../models"
)

type UsersDAO struct {
	Server 		string
	Database	string 
}

const (
	COLLECTION =  "users"
)

