package main

import (
	"context"
	"log"
	"net"
  "os"
  "fmt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	healthpb "google.golang.org/grpc/health/grpc_health_v1"

	"github.com/hanapedia/microservice-topologies/chain_example/chain_3/connections/mongo"
	pb "github.com/hanapedia/microservice-topologies/chain_example/chain_3/pb_chain_3"
)

type chain_3Server struct {
	pb.UnimplementedChain_3Server
}

const (
  ListenPort = "3001"
  ChainNextAddress = "localhost:3002"
  DbAddress = "mongodb://localhost:27017"
  DbName = "chain"
)

var mongoConnection *mongo.Mongo


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

  lis, err := net.Listen("tcp", fmt.Sprintf(":%s", listenPort))
  if err != nil {
    log.Fatalf("failed to listen: %v", err)
  }
  var opts []grpc.ServerOption
  grpcServer := grpc.NewServer(opts...)
  svc := new(chain_3Server)
  pb.RegisterChain_3Server(grpcServer, svc)
  healthpb.RegisterHealthServer(grpcServer, svc)

  mongoConnection, err = mongo.InitMongo(dbAddress, dbName, "chain1")
  if err != nil {
    log.Fatalf("failed to connect to mongodb: %v", err)
  }
  log.Printf("Established connection with %v", dbAddress)
  defer mongoConnection.Disconnect()

  
  log.Printf("Starting server at port %v", listenPort)
  grpcServer.Serve(lis)
}

func (s chain_3Server) GetIds(ctx context.Context, req *pb.Req) (*pb.Res, error) {
  
  res := new(pb.Res)
  res.Ids = req.Ids

  newId, err := mongoConnection.GetItem(req.Ids[len(req.Ids)-1])
  if err == nil {
      res.Ids = append(res.Ids, newId)
    } else {
    log.Fatalf("Failed to retrieve item: %v", err)
  }

  	return res, err
}

func (s *chain_3Server) Check(ctx context.Context, req *healthpb.HealthCheckRequest) (*healthpb.HealthCheckResponse, error) {
	return &healthpb.HealthCheckResponse{Status: healthpb.HealthCheckResponse_SERVING}, nil
}

func (s *chain_3Server) Watch(req *healthpb.HealthCheckRequest, ws healthpb.Health_WatchServer) error {
	return status.Errorf(codes.Unimplemented, "health check via Watch not implemented")
}
