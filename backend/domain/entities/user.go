package entities

import (
    "time"
    "go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
    ID            primitive.ObjectID `json:"-" bson:"_id,omitempty"`
    Username      string             `bson:"username"`
    Email         string             `bson:"email"`
    Password      string             `bson:"password"`
    Role          string             `bson:"role"`
    CreatedAt     time.Time          `bson:"createdAt"`
    UpdatedAt     time.Time          `bson:"updatedAt"`
}




