package repository

import (
	"backend-starter-project/domain/entities"
	"backend-starter-project/domain/interfaces"
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type defaultImageRepository struct {
	db *mongo.Database
	collection *mongo.Collection
}

func NewDefaultImageRepository(db *mongo.Database) interfaces.DefaultImageRepository {
	return &defaultImageRepository{db: db, collection: db.Collection("default_image")}
}

func (r *defaultImageRepository) CreateDefaultImage(url,publicID string) error {
	_,err:=r.collection.InsertOne(context.TODO(),bson.D{{"url",url},{"publicID",publicID}})
	if err!=nil{
		return err
	}
	return nil

}

func (r *defaultImageRepository) DeleteDefaultImage(id string) error {
	imageId,err:=primitive.ObjectIDFromHex(id)
	if err!=nil{
		return err
	}
	_,err=r.collection.DeleteOne(context.TODO(),bson.D{{"_id",imageId}})
	if err!=nil{
		return err
	}
	return nil
}
func (r *defaultImageRepository) GetDefaultImage(id string) (*entities.DefaultImage,error) {
	imageId,err:=primitive.ObjectIDFromHex(id)
	if err!=nil{
		return nil,err
	}
	var defaultImage entities.DefaultImage
	err=r.collection.FindOne(context.TODO(),bson.D{{"_id",imageId}}).Decode(&defaultImage)
	if err!=nil{
		return nil,err
	}
	return &defaultImage,nil
}	

func (r *defaultImageRepository) GetDefaultImages() ([]*entities.DefaultImage,error) {
	cursor,err:=r.collection.Find(context.TODO(),bson.D{})
	if err!=nil{
		return nil,err
	}
	var defaultImages []*entities.DefaultImage
	cursor.All(context.TODO(),&defaultImages)
	return defaultImages,nil
}