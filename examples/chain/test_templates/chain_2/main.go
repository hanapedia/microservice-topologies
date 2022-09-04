package main

import (
	"context"
	"log"
	"net"

	"github.com/hanapedia/microservice-topologies/chain/chain_2/connections/mongo"
	pb "github.com/hanapedia/microservice-topologies/chain/chain_2/pb"
	"google.golang.org/grpc"
)

type chain_2Server struct {
	pb.UnimplementedChian_1Server
}

var mongoConnection *mongo.Mongo

func (s chain_2Server) GetIds(ctx context.Context, req *pb.Req) (*pb.Res, error) {
  newId, err := mongoConnection.GetItem(req.Ids[len(req.Ids)-1])
  if err != nil {
    log.Fatalf("Failed to retrieve item: %v", err)
  }
	returnArr := req.Ids
  returnArr = append(returnArr, newId)
	// Implement db read logic
  res := pb.Res {
    Ids: returnArr,
  }

	return &res, nil
}

func newServer() chain_2Server {
  return chain_2Server{}
}

func main() {
  lis, err := net.Listen("tcp", "localhost:3000")
  if err != nil {
    log.Fatalf("failed to listen: %v", err)
  }

  var opts []grpc.ServerOption

  grpcServer := grpc.NewServer(opts...)
  pb.RegisterChian_1Server(grpcServer, newServer())

  mongoConnection, err = mongo.InitMongo("mongodb://localhost:27017", "chain", "chain1")
  if err != nil {
    log.Fatalf("failed to connect to mongodb: %v", err)
  }
  defer mongoConnection.Disconnect()

  grpcServer.Serve(lis)
}
