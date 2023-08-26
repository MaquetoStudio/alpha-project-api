package repositories

import (
	"context"

	"alpha-project-api/database"
	"alpha-project-api/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepo struct {
	Collection *mongo.Collection
}

func NewUserRepo() *UserRepo {
	return &UserRepo{
		Collection: database.NewDbInstance().Collection("users"), // Récupère la connexion singleton à la base de données
	}
}

func (r *UserRepo) Insert(user models.User) (*models.User, error) {
	insertResult, err := r.Collection.InsertOne(context.TODO(), user)
	if err != nil {
		return nil, err
	}

	insertedID := insertResult.InsertedID.(primitive.ObjectID).Hex()

	user.ID = insertedID
	return &user, nil
}

func (r *UserRepo) GetUserByID(id string) (*models.User, error) {
	var result models.User
	err := r.Collection.FindOne(context.TODO(), bson.M{"id": id}).Decode(result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
