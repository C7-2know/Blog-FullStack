package entities

import "go.mongodb.org/mongo-driver/bson/primitive"

type DefaultImage struct {
	Id  primitive.ObjectID `json:"_id" bson:"_id"`
	PublicID string `json:"publicID" bson:"publicID"`
	Url string `json:"url" bson:"url"`
}