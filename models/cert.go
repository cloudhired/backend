package models

type CertProfile struct {
	Name     string `bson:"name" json:"name"`
	Icon     string `bson:"icon" json:"icon"`
	IssuedBy string `bson:"issued_by" json:"issued_by"`
}
