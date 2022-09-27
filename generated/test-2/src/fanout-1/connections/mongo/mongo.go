package mongo

import (
	"context"
	"errors"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type Mongo struct {
	Client     *mongo.Client
	Collection *mongo.Collection
}

type Schema struct {
	Key   int `bson:"key"`
	Value int `bson:"value"`
}

func InitMongo(uri string, dbUser string, dbPassword string, dbName string, collectionName string) (*Mongo, error) {
	ctx := context.Background()
	credential := options.Credential{
		Username: dbUser,
		Password: dbPassword,
	}
	clientOpts := options.Client().ApplyURI(uri).SetAuth(credential)
	client, err := mongo.Connect(ctx, clientOpts)
	if err != nil {
		return &Mongo{}, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		log.Printf("cannot reach mongo db server: %v", err)
	}
	defer cancel()

	collection := client.Database(dbName).Collection(collectionName)

	return &Mongo{Client: client, Collection: collection}, nil
}

func (m *Mongo) GetItem(id int32) (int32, error) {
	ctx := context.Background()
	filter := bson.D{primitive.E{Key: "key", Value: id}}
	var result Schema

	err := m.Collection.FindOne(ctx, filter).Decode(&result)
	if err == mongo.ErrNoDocuments {
		return 0, errors.New("No document found with matching id")
	}
	if err != nil {
		return 0, err
	}
	return int32(result.Value), nil
}

func (m *Mongo) Disconnect() {
	if err := m.Client.Disconnect(context.Background()); err != nil {
		panic(err)
	}
}
