package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	Username        primitive.ObjectID `bson:"_id" json:"id"`
	FullName        string             `bson:"fullname" json:"fullname"`
	FirstName       string             `bson:"fname" json:"fname"`
	LastName        string             `bson:"lname" json:"last_name"`
	JobTitle        string             `bson:"job_title" json:"job_title"`
	CurrentLocation string             `bson:"current_loc" json:"current_location"`
	Company         string             `bson:"company" json:"company"`
	YOE             string             `bson:"yoe" json:"yoe"`
	PersonalSite    string             `bson:"personal_site" json:"personal_site"`
	LinkedinHandle  string             `bson:"linkedin_handle" json:"linkedin_handle"`
	GithubHandle    string             `bson:"github_handle" json:"github_handle"`
	Email           string             `bson:"email" json:"email"`
	Skills          []string           `bson:"skills" json:"skills"`
	Certs           []Cert             `bson:"certs" json:"certs"`
	GHContributions string             `bson:"gh_contributions" json:"gh_contributions"`
	Portfolios      []Portfolio        `bson:"portfolios" json:"portfolios"`
	Courses         []Course           `bson:"courses" json:"courses"`
	Intro           string             `bson:"intro" json:"intro"`
}

type Cert struct {
	CertName    string             `bson: "cert_name" json:"cert_name"`
	DateIssued  primitive.DateTime `bson:"date_issued" json:"date_issued"`
	DateExpired primitive.DateTime `bson:"date_expired" json:"date_expired"`
	VerifyLink  string             `bson:"verify_link" json:"verify_link"`
}

type Portfolio struct {
	Pics        string `bson:"pics" json:"pics"`
	Description string `bson:"description" json:"description"`
}

type Course struct {
	Name         string             `bson:"name" json:"name"`
	Link         string             `bson:"link" json:"link"`
	DateFinished primitive.DateTime `bson:"date_finished" json:"date_finished"`
	Outcome      string             `bson:"outcome" json:"outcome"`
}
