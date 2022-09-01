package main

import (
	"context"
	"log"
	"net"

	"github.com/hanapedia/microservice-topologies/chain/chain_1/connections/mongo"
	pb "github.com/hanapedia/microservice-topologies/chain/chain_1/pb_chain_1"
	pbc "github.com/hanapedia/microservice-topologies/chain/chain_1/pb_chain_2"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	healthpb "google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/internal/status"
)

type chain_1Server struct {
	pb.UnimplementedChain_1Server
}

var mongoConnection *mongo.Mongo

var nextChain pbc.Chain_2Client

func (s chain_1Server) GetIds(ctx context.Context, req *pb.Req) (*pb.Res, error) {
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

func (s *chain_1Server) Check(ctx context.Context, req *healthpb.HealthCheckRequest) (*healthpb.HealthCheckResponse, error) {
	return &healthpb.HealthCheckResponse{Status: healthpb.HealthCheckResponse_SERVING}, nil
}

func (s *chain_1Server) Watch(req *healthpb.HealthCheckRequest, ws healthpb.Health_WatchServer) error {
	return status.Errorf(codes.Unimplemented, "health check via Watch not implemented")
}

func newServer() chain_1Server {
  return chain_1Server{}
}

func main() {
  lis, err := net.Listen("tcp", "localhost:3000")
  if err != nil {
    log.Fatalf("failed to listen: %v", err)
  }
  var opts []grpc.ServerOption
  grpcServer := grpc.NewServer(opts...)
  svc := newServer()
  pb.RegisterChain_1Server(grpcServer, svc)
  healthpb.RegisterHealthServer(grpcServer, &svc)

  mongoConnection, err = mongo.InitMongo("mongodb://localhost:27017", "chain", "chain1")
  if err != nil {
    log.Fatalf("failed to connect to mongodb: %v", err)
  }
  defer mongoConnection.Disconnect()

  var optsClient []grpc.DialOption
  optsClient = append(optsClient, grpc.WithTransportCredentials(insecure.NewCredentials()))

  conn, err := grpc.Dial("localhost:3002", optsClient...)
  if err != nil {
    log.Printf("Cannot establish connection with the server: %v", err)
  }
  defer conn.Close()

  nextChain = pbc.NewChain_2Client(conn)

  grpcServer.Serve(lis)
}
