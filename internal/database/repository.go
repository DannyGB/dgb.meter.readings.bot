package database

import (
	"context"
	"dgb/meter.readings.bot/internal/configuration"
	"errors"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client

type Repository struct {
	config configuration.Configuration
}

func (repository *Repository) GetLatest() []bson.M {
	connect(repository.config)
	coll := repository.getCollection()

	myOptions := options.Find()
	myOptions.
		SetSort(bson.M{"$natural": -1}).
		SetLimit(2).
		SetProjection(bson.D{{Key: "reading", Value: 1}, {Key: "rate", Value: 1}, {Key: "_id", Value: 0}})

	cursor, err := coll.Find(context.TODO(), bson.M{}, myOptions)

	if err == mongo.ErrNoDocuments {
		return nil
	}

	var results []bson.M
	if err = cursor.All(context.TODO(), &results); err != nil {
		panic(err)
	}

	return results
}

func (repository *Repository) Insert(data bson.M) (id interface{}, err error) {

	connect(repository.config)
	coll := repository.getCollection()

	date, _ := time.Parse(time.RFC3339Nano, data["readingdate"].(string))

	data["readingdate"] = date

	result, err := coll.InsertOne(context.TODO(), data)

	if err != nil {
		return nil, errors.New("Could not insert document")
	}

	return result.InsertedID, nil
}

func connect(config configuration.Configuration) {

	if client != nil {
		return
	}

	var err error
	client, err = mongo.Connect(context.TODO(), options.Client().ApplyURI(config.MONGO_CONNECTION))

	if err != nil {
		panic(err)
	}
}

func (repository *Repository) getCollection() *mongo.Collection {
	return client.Database(repository.config.MONGO_DB).Collection(repository.config.MONGO_COLLECTION)
}

func NewRepository(cfg configuration.Configuration) *Repository {
	return &Repository{
		cfg,
	}
}
