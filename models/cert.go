package models

type CertProfile struct {
	Name     string `bson:"name" json:"name"`
	Id       string `bson:"id" json:"id"`
	IssuedBy string `bson:"issued_by" json:"issued_by"`
}
