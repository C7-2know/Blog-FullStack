package repository

import (
	"backend-starter-project/domain/entities"
	"backend-starter-project/domain/interfaces"
	"context"
	"errors"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type profileRepository struct {
	db         *mongo.Database
	collection *mongo.Collection
	context    context.Context
}

func NewProfileRepository(ctx context.Context, db *mongo.Database) interfaces.ProfileRepository {
	return profileRepository{db: db, collection: db.Collection("profile"), context: ctx}
}
func (repo profileRepository) GetProfiles() ([]*entities.Profile, error) {
	cursor, err := repo.collection.Find(context.Background(), bson.D{})
	if err != nil {
		return nil, err
	}
	var profiles []*entities.Profile
	cursor.All(context.Background(), &profiles)
	return profiles, nil
}

func (repo profileRepository) GetUserProfile(user_id string) (*entities.Profile, error) {
	id, err := primitive.ObjectIDFromHex(user_id)
	if err != nil {
		return nil, err
	}
	filter := bson.D{{"userId", id}}
	fmt.Println(user_id,id)
	user := repo.collection.FindOne(context.TODO(), filter)
	if user.Err() != nil {
		fmt.Println("err", user.Err())
		return nil, errors.New("couldn't get the profile")
	}
	var profile entities.Profile
	err=user.Decode(&profile)
	if err!=nil{
		return nil,err
	}
	fmt.Println("profile from repo", profile)
	return &profile, nil

}

func (repo profileRepository) CreateUserProfile(profile *entities.Profile) (*entities.Profile, error) {
	_, err := repo.collection.InsertOne(repo.context, &profile)
	if err != nil {
		return nil, err
	}
	return profile, nil
}

func (repo profileRepository) UpdateUserProfile(profile *entities.Profile) (*entities.Profile, error) {
	user_id := profile.UserID
	filter := bson.D{{"userId", user_id}}
	_, err := repo.collection.UpdateOne(context.TODO(), filter, profile)
	if err != nil {
		return nil, err
	}

	return profile, nil
}

func (repo profileRepository) DeleteUserProfile(user_id string) error {
	filter := bson.D{{"userId", user_id}}
	_, err := repo.collection.DeleteOne(repo.context, filter)
	if err != nil {
		return err
	}

	return nil
}
