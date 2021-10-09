package models

import (
	"gopkg.in/mgo.v2/bson"
	"github.com/path/to/timestamp"
)

type User struct{
	Id bson.ObjectId `json: "id" bson:"_id"'`
	Name string    `json:"name" bson:"name"`
	Email string	`json:"email" bson:"email"`
	Password string	`json:"password" bson:"password"`
}

type Post struct{
	Id bson.ObjectId `json: "id" bson:"_id"'`
	Caption string    `json:"name" bson:"name"`
	Image_URL string	`json:"image" bson:"image"`
	Created_at  *timestamp.Timestamp  `json:"created_at,omitempty" bson:"created_at,omitempty" `
}

