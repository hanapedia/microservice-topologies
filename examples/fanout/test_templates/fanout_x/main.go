package main

import (
	"context"
	"log"
	"net"
	"os"

	// "github.com/hanapedia/microservice-topologies/fanout/fanout_x/connections/mongo"
	// pb "github.com/hanapedia/microservice-topologies/fanout/fanout_x/pb_fanout_x"

	"github.com/hanapedia/microservice-topologies/fanout/test_templates/fanout_x/connections/mongo"
	pb "github.com/hanapedia/microservice-topologies/fanout/test_templates/fanout_x/pb_fanout_x"
	"google.golang.org/grpc"
)

type fanout_xServer struct {
	pb.UnimplementedFanoutXServer
}

var mongoConnection *mongo.Mongo

const (
	ListenPort = "3002"
	DbAddress  = "mongodb://localhost:27017"
	DbName     = "fanout"
)

func (s fanout_xServer) GetIds(ctx context.Context, req *pb.Req) (*pb.Res, error) {
	newId, err := mongoConnection.GetItem(req.Ids[len(req.Ids)-1])
	if err != nil {
		log.Fatalf("Failed to retrieve item: %v", err)
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
	dbName := DbName
	if os.Getenv("DB_NAME") != "" {
		dbName = os.Getenv("DB_NAME")
	}

	svc := new(fanout_xServer)
	lis, err := net.Listen("tcp", listenPort)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	var opts []grpc.ServerOption

	grpcServer := grpc.NewServer(opts...)
	pb.RegisterFanoutXServer(grpcServer, svc)

	mongoConnection, err = mongo.InitMongo(dbAddress, dbName, "fanout_x")
	if err != nil {
		log.Fatalf("failed to connect to mongodb: %v", err)
	}
	defer mongoConnection.Disconnect()

	grpcServer.Serve(lis)
}
