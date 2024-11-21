package dtos

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type RegisterUserDto struct {
	ID       primitive.ObjectID `json:"-"` 
	Username string             `json:"username" binding:"required"`
	Email    string             `json:"email" binding:"required"`
	Password string             `json:"password" binding:"required"`
}

type UserLoginDto struct{
	Email string `json:"email"`
	Password string `json:"password"`
}