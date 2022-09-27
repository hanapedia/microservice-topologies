package main

import (
	"context"
	"log"
	"net"
	"os"
  "fmt"

	"github.com/hanapedia/microservice-topologies/test-2/fanout-3/connections/mongo"
	pb "github.com/hanapedia/microservice-topologies/test-2/fanout-3/pb-fanout-3"
	"google.golang.org/grpc"
)

type fanout_3Server struct {
	pb.UnimplementedFanout_3Server
}

var mongoConnection *mongo.Mongo

const (
	ListenPort = "4004"
	DbAddress  = "mongodb://localhost:27017"
  DbUser = "root"
  DbPassword = "example"
  DbName = "test-2"
  CollectionName = "fanout3"
)

func (s fanout_3Server) GetIds(ctx context.Context, req *pb.Req) (*pb.Res, error) {
	newId, err := mongoConnection.GetItem(req.Ids[len(req.Ids)-1])
	if err != nil {
		log.Printf("Failed to retrieve item at fanout_3: %v", err)
	}
	returnArr := req.Ids
	returnArr = append(returnArr, newId)
	// Implement db read logic
	res := pb.Res{
		Ids: returnArr,
	}

	return &res, nil
}

func main() {
	listenPort := ListenPort
	if os.Getenv("LISTEN_PORT") != "" {
		listenPort = os.Getenv("LISTEN_PORT")
	}
	dbAddress := DbAddress
	if os.Getenv("DB_ADDRESS") != "" {
		dbAddress = os.Getenv("DB_ADDRESS")
	}
	dbUser := DbUser
	if os.Getenv("DB_USER") != "" {
		dbUser = os.Getenv("DB_USER")
	}
	dbPassword := DbPassword
	if os.Getenv("DB_PASSWORD") != "" {
		dbPassword = os.Getenv("DB_PASSWORD")
	}
	dbName := DbName
	if os.Getenv("DB_NAME") != "" {
		dbName = os.Getenv("DB_NAME")
	}
	collectionName := CollectionName
	if os.Getenv("COLLECTION_NAME") != "" {
		collectionName = os.Getenv("COLLECTION_NAME")
	}

	svc := new(fanout_3Server)
  lis, err := net.Listen("tcp", fmt.Sprintf(":%s", listenPort))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	var opts []grpc.ServerOption

	grpcServer := grpc.NewServer(opts...)
	pb.RegisterFanout_3Server(grpcServer, svc)

	mongoConnection, err = mongo.InitMongo(dbAddress, dbUser, dbPassword, dbName, collectionName)
	if err != nil {
		log.Fatalf("failed to connect to mongodb: %v", err)
	}
	defer mongoConnection.Disconnect()

	grpcServer.Serve(lis)
}
