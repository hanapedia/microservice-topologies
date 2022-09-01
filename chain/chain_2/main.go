package main

import (
	"context"
	"log"
	"net"

	"github.com/hanapedia/microservice-topologies/chain/chain_2/connections/mongo"
	pb "github.com/hanapedia/microservice-topologies/chain/chain_2/pb_chain_2"
	pbc "github.com/hanapedia/microservice-topologies/chain/chain_2/pb_chain_3"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type chain_2Server struct {
	pb.UnimplementedChain_2Server
}

var mongoConnection *mongo.Mongo

var nextChain pbc.Chain_3Client

func (s chain_2Server) GetIds(ctx context.Context, req *pb.Req) (*pb.Res, error) {
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

func newServer() chain_2Server {
  return chain_2Server{}
}

func main() {
  lis, err := net.Listen("tcp", "localhost:3002")
  if err != nil {
    log.Fatalf("failed to listen: %v", err)
  }
  var opts []grpc.ServerOption
  grpcServer := grpc.NewServer(opts...)
  pb.RegisterChain_2Server(grpcServer, newServer())

  mongoConnection, err = mongo.InitMongo("mongodb://localhost:27017", "chain", "chain1")
  if err != nil {
    log.Fatalf("failed to connect to mongodb: %v", err)
  }
  defer mongoConnection.Disconnect()

  var optsClient []grpc.DialOption
  optsClient = append(optsClient, grpc.WithTransportCredentials(insecure.NewCredentials()))

    conn, err := grpc.Dial("localhost:3003", optsClient...)
  if err != nil {
    log.Printf("Cannot establish connection with the server: %v", err)
  }
  defer conn.Close()

  nextChain = pbc.NewChain_3Client(conn)
  
  grpcServer.Serve(lis)
}
