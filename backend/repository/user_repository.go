package repository

import (
	"backend-starter-project/domain/dtos"
	"backend-starter-project/domain/entities"
	"backend-starter-project/domain/interfaces"
	"context"
	"errors"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type userRepository struct {
	collection *mongo.Collection
}

func NewUserRepository(collection *mongo.Collection) interfaces.UserRepository {
	return &userRepository{
		collection: collection,
	}
}

func (repo *userRepository) CreateUser(user *dtos.RegisterUserDto) (*entities.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	existedUser, err := repo.FindUserByEmail(user.Email)
	if existedUser != nil {
		return nil, errors.New("User already exists")
	}

	newUser:=entities.User{}
	newUser.ID=primitive.NewObjectID()
	newUser.Username=user.Username
	newUser.Email=user.Email
	newUser.Password=user.Password
	newUser.Role="user"
	_, err = repo.collection.InsertOne(ctx, newUser)
	if err != nil {
		return nil, err
	}

	return &newUser, nil
}

func (repo *userRepository) FindUserByEmail(email string) (*entities.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var user entities.User
	err := repo.collection.FindOne(ctx, bson.M{"email": email}).Decode(&user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (repo *userRepository) DeleteUser(userId string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := repo.collection.DeleteOne(ctx, bson.M{"_id": userId})
	if err != nil {
		return err
	}

	return nil
}

func (repo *userRepository) FindUserById(userId string) (*entities.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var user entities.User
	userID, err := primitive.ObjectIDFromHex(userId)
	err = repo.collection.FindOne(ctx, bson.M{"_id": userID}).Decode(&user)
	if err != nil {
		// to do
		// handle error type
		return nil, err
	}

	return &user, nil

}

func (repo *userRepository) UpdateUser(user *entities.User) (*entities.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := repo.collection.UpdateOne(ctx, bson.M{"_id": user.ID}, bson.M{"$set": user})
	if err != nil {
		return nil, err
	}

	return user, nil

}
