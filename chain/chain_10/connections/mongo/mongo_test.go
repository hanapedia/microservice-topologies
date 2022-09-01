package mongo

import (
	"context"
	"log"
	"testing"

	"go.mongodb.org/mongo-driver/bson"
)

func TestMongoConnection(t *testing.T) {
  MongoConnection, err := InitMongo("mongodb://localhost:27017", "chain", "chain1")
  if err != nil {
    t.Errorf("Connection failed: %v", err)
  }
  // i, err := MongoConnection.GetItem(3)
  // if err != nil {
  //   t.Errorf("data retrieval failed: %v", err)
  // }
  // log.Println(i)

  // res, err := MongoConnection.client.Database("chain").ListCollectionNames(
  //   context.Background(),
  //   bson.D{},
  // )
  result := &Schema{}
  filter := bson.D{{ "key", 3 }}
  err = MongoConnection.Collection.FindOne(context.Background(), filter).Decode(result)
  if err!= nil {
    log.Fatal(err)
  }
  log.Println(result.Value)
}
