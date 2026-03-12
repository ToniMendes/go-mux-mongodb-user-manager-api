package mongodb

import (
	"context"
	"go-mux-mongodb-user-manager-api/internal/domain"
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type MongoRepository struct {
	Collection *mongo.Collection
}

func NewMongoRepository(col *mongo.Collection) *MongoRepository {
	return &MongoRepository{
		Collection: col,
	}
}

func IndexUnique(col *mongo.Collection) error {
	model := mongo.IndexModel{
		Keys:    bson.D{{Key: "email", Value: 1}},
		Options: options.Index().SetUnique(true),
	}
	_, err := col.Indexes().CreateOne(context.Background(), model)
	return err
}

func (repo *MongoRepository) Create(model *domain.User) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	doc := bson.M{
		"name":     model.Name,
		"email":    model.Email,
		"password": model.PasswordHash,
	}

	_, err := repo.Collection.InsertOne(ctx, doc)
	if err != nil {
		if mongo.IsDuplicateKeyError(err) {
			return err
		}
		return err
	}

	return nil
}

func (repo *MongoRepository) GetAll() ([]domain.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	cursor, err := repo.Collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}

	var result []domain.User
	if err = cursor.All(ctx, &result); err != nil {
		return nil, err
	}

	return result, nil
}

func (repo *MongoRepository) GetByEmail(email string) (*domain.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	filter := bson.M{"email": email}

	var model *domain.User

	err := repo.Collection.FindOne(ctx, filter).Decode(&model)
	if err != nil {
		return nil, err
	}

	return model, nil
}

func (repo *MongoRepository) UpdateName(name, email string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	filter := bson.M{"email": email}
	doc := bson.M{"$set": bson.M{"name": name}}

	_, err := repo.Collection.UpdateOne(ctx, filter, doc)
	if err != nil {
		return err
	}

	return nil
}

func (repo *MongoRepository) UpdateEmail(newEmail, email string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	filter := bson.M{"email": email}
	doc := bson.M{"$set": bson.M{"email": newEmail}}

	_, err := repo.Collection.UpdateOne(ctx, filter, doc)
	if err != nil {
		return err
	}

	return nil
}
