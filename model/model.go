package model

import "go.mongodb.org/mongo-driver/bson/primitive"

// Cloth struct type
type Cloth struct {
	ID    primitive.ObjectID `bson:"_id"`
	Src   string             `bson:"src"`
	Name  string             `bson:"name"`
	Brand string             `bson:"brand"`
	Site  string             `bson:"site"`
	Price int                `bson:"price"`
	Img   string             `bson:"img"`
}

// DBinfo struct type
type DBinfo struct {
	Username string
	Password string
	Hostname string
	Port     string
}
