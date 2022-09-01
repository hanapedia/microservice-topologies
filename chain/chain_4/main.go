package main

import (
	"context"
	"log"
	"net"

	"github.com/hanapedia/microservice-topologies/chain/chain_4/connections/mongo"
	pb "github.com/hanapedia/microservice-topologies/chain/chain_4/pb_chain_4"
	pbc "github.com/hanapedia/microservice-topologies/chain/chain_4/pb_chain_5"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type chain_4Server struct {
	pb.UnimplementedChain_4Server
}

var mongoConnection *mongo.Mongo

var nextChain pbc.Chain_5Client

func (s chain_4Server) GetIds(ctx context.Context, req *pb.Req) (*pb.Res, error) {
  newId, err := mongoConnection.GetItem(req.Ids[len(req.Ids)-1])
  if err != nil {
    log.Fatalf("Failed to retrieve item: %v", err)
  }
	returnArr := req.Ids
  returnArr = append(returnArr, newId)
	// Implement db read logic
    next := pbc.Req {
    Ids: returnArr,
  }

  nextRes, err := nextChain.GetIds(context.Background(), &next)
  if err != nil {
    log.Fatalf("Failed to retrieve ids: %v", err)
  }

  res := pb.Res {
    Ids: nextRes.Ids,
  }
  
	return &res, nil
}

func newServer() chain_4Server {
  return chain_4Server{}
}

func main() {
  lis, err := net.Listen("tcp", "localhost:3004")
  if err != nil {
    log.Fatalf("failed to listen: %v", err)
  }
  var opts []grpc.ServerOption
  grpcServer := grpc.NewServer(opts...)
  pb.RegisterChain_4Server(grpcServer, newServer())

  mongoConnection, err = mongo.InitMongo("mongodb://localhost:27017", "chain", "chain1")
  if err != nil {
    log.Fatalf("failed to connect to mongodb: %v", err)
  }
  defer mongoConnection.Disconnect()

  var optsClient []grpc.DialOption
  optsClient = append(optsClient, grpc.WithTransportCredentials(insecure.NewCredentials()))

    conn, err := grpc.Dial("localhost:3005", optsClient...)
  if err != nil {
    log.Printf("Cannot establish connection with the server: %v", err)
  }
  defer conn.Close()

  nextChain = pbc.NewChain_5Client(conn)
  
  grpcServer.Serve(lis)
}
