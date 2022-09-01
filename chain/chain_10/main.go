package main

import (
	"context"
	"log"
	"net"

	"github.com/hanapedia/microservice-topologies/chain/chain_10/connections/mongo"
	pb "github.com/hanapedia/microservice-topologies/chain/chain_10/pb_chain_10"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type chain_10Server struct {
	pb.UnimplementedChain_10Server
}

var mongoConnection *mongo.Mongo


func (s chain_10Server) GetIds(ctx context.Context, req *pb.Req) (*pb.Res, error) {
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

func newServer() chain_10Server {
  return chain_10Server{}
}

func main() {
  lis, err := net.Listen("tcp", "localhost:3010")
  if err != nil {
    log.Fatalf("failed to listen: %v", err)
  }
  var opts []grpc.ServerOption
  grpcServer := grpc.NewServer(opts...)
  pb.RegisterChain_10Server(grpcServer, newServer())

  mongoConnection, err = mongo.InitMongo("mongodb://localhost:27017", "chain", "chain1")
  if err != nil {
    log.Fatalf("failed to connect to mongodb: %v", err)
  }
  defer mongoConnection.Disconnect()

  var optsClient []grpc.DialOption
  optsClient = append(optsClient, grpc.WithTransportCredentials(insecure.NewCredentials()))

  
  grpcServer.Serve(lis)
}
