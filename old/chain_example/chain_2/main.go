package main

import (
	"context"
	"log"
	"net"
  "os"
  "fmt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"
	healthpb "google.golang.org/grpc/health/grpc_health_v1"

	"github.com/hanapedia/microservice-topologies/chain_example/chain_2/connections/mongo"
	pb "github.com/hanapedia/microservice-topologies/chain_example/chain_2/pb_chain_2"
	pbc "github.com/hanapedia/microservice-topologies/chain_example/chain_2/pb_chain_3"
)

type chain_2Server struct {
	pb.UnimplementedChain_2Server
}

const (
  ListenPort = "3001"
  ChainNextAddress = "localhost:3002"
  DbAddress = "mongodb://localhost:27017"
  DbName = "chain"
)

var mongoConnection *mongo.Mongo

var nextChain pbc.Chain_3Client

func main() {
	listenPort := ListenPort
	if os.Getenv("LISTEN_PORT") != "" {
		listenPort = os.Getenv("LISTEN_PORT")
	}
  	chainNextAddress := ChainNextAddress
	if os.Getenv("CHAIN_NEXT_ADDRESS") != "" {
		chainNextAddress = os.Getenv("CHAIN_NEXT_ADDRESS")
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
  svc := new(chain_2Server)
  pb.RegisterChain_2Server(grpcServer, svc)
  healthpb.RegisterHealthServer(grpcServer, svc)

  mongoConnection, err = mongo.InitMongo(dbAddress, dbName, "chain1")
  if err != nil {
    log.Fatalf("failed to connect to mongodb: %v", err)
  }
  log.Printf("Established connection with %v", dbAddress)
  defer mongoConnection.Disconnect()

    var optsClient []grpc.DialOption
  optsClient = append(optsClient, grpc.WithTransportCredentials(insecure.NewCredentials()))

  conn, err := grpc.Dial(chainNextAddress, optsClient...)
  if err != nil {
    log.Printf("Cannot establish connection with the server: %v", err)
  }
  log.Printf("Dialed and established connection with %v", chainNextAddress)
  defer conn.Close()

  nextChain = pbc.NewChain_3Client(conn)
  
  log.Printf("Starting server at port %v", listenPort)
  grpcServer.Serve(lis)
}

func (s chain_2Server) GetIds(ctx context.Context, req *pb.Req) (*pb.Res, error) {
    next := new(pbc.Req)
  next.Ids = req.Ids
  
  res := new(pb.Res)
  res.Ids = req.Ids

  newId, err := mongoConnection.GetItem(req.Ids[len(req.Ids)-1])
  if err == nil {
      next.Ids = append(next.Ids, newId)
    } else {
    log.Fatalf("Failed to retrieve item: %v", err)
  }

    nextRes, err := nextChain.GetIds(context.Background(), next)
  if err == nil {
    res.Ids = nextRes.Ids
  } else {
    log.Fatalf("Failed to retrieve ids: %v", err)
  }
  	return res, err
}

func (s *chain_2Server) Check(ctx context.Context, req *healthpb.HealthCheckRequest) (*healthpb.HealthCheckResponse, error) {
	return &healthpb.HealthCheckResponse{Status: healthpb.HealthCheckResponse_SERVING}, nil
}

func (s *chain_2Server) Watch(req *healthpb.HealthCheckRequest, ws healthpb.Health_WatchServer) error {
	return status.Errorf(codes.Unimplemented, "health check via Watch not implemented")
}
