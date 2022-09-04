package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	healthpb "google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/status"

	pb "github.com/hanapedia/microservice-topologies/fanout_example/fanout_0/pb_fanout_0"
	pb1 "github.com/hanapedia/microservice-topologies/fanout_example/fanout_0/pb_fanout_1"
	pb2 "github.com/hanapedia/microservice-topologies/fanout_example/fanout_0/pb_fanout_2"
	pb3 "github.com/hanapedia/microservice-topologies/fanout_example/fanout_0/pb_fanout_3"
)

type fanout_0Server struct {
	pb.UnimplementedFanout_0Server
}

const (
	ListenPort           = "3001"
	FanoutClient1Address = "localhost:4002"
	FanoutClient2Address = "localhost:4003"
	FanoutClient3Address = "localhost:4004"
)

var fanout_1Client pb1.Fanout_1Client
var fanout_2Client pb2.Fanout_2Client
var fanout_3Client pb3.Fanout_3Client

func main() {
	listenPort := ListenPort
	if os.Getenv("LISTEN_PORT") != "" {
		listenPort = os.Getenv("LISTEN_PORT")
	}
	fanoutClient1Address := FanoutClient1Address
	if os.Getenv("FANOUT_CLIENT_1_ADDRESS") != "" {
		fanoutClient1Address = os.Getenv("FANOUT_CLIENT_1_ADDRESS")
	}
	fanoutClient2Address := FanoutClient2Address
	if os.Getenv("FANOUT_CLIENT_2_ADDRESS") != "" {
		fanoutClient2Address = os.Getenv("FANOUT_CLIENT_2_ADDRESS")
	}
	fanoutClient3Address := FanoutClient3Address
	if os.Getenv("FANOUT_CLIENT_3_ADDRESS") != "" {
		fanoutClient3Address = os.Getenv("FANOUT_CLIENT_3_ADDRESS")
	}

	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", listenPort))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	svc := new(fanout_0Server)
	pb.RegisterFanout_0Server(grpcServer, svc)
	healthpb.RegisterHealthServer(grpcServer, svc)

	var optsClient []grpc.DialOption
	optsClient = append(optsClient, grpc.WithTransportCredentials(insecure.NewCredentials()))

	conn_1, err := grpc.Dial(fanoutClient1Address, optsClient...)
	if err != nil {
		log.Fatalf("Cannot establish connection with the server: %v", err)
	}
	log.Printf("Dialed and established connection with %v", fanoutClient1Address)
	defer conn_1.Close()

	fanout_1Client = pb1.NewFanout_1Client(conn_1)
	conn_2, err := grpc.Dial(fanoutClient2Address, optsClient...)
	if err != nil {
		log.Fatalf("Cannot establish connection with the server: %v", err)
	}
	log.Printf("Dialed and established connection with %v", fanoutClient2Address)
	defer conn_2.Close()

	fanout_2Client = pb2.NewFanout_2Client(conn_2)
	conn_3, err := grpc.Dial(fanoutClient3Address, optsClient...)
	if err != nil {
		log.Fatalf("Cannot establish connection with the server: %v", err)
	}
	log.Printf("Dialed and established connection with %v", fanoutClient3Address)
	defer conn_3.Close()

	fanout_3Client = pb3.NewFanout_3Client(conn_3)

	log.Printf("Starting server at port %v", listenPort)
	grpcServer.Serve(lis)
}

func (s fanout_0Server) GetIds(ctx context.Context, req *pb.Req) (*pb.Res, error) {
	res := new(pb.Res)
	res.Ids = req.Ids

  fanout_1Req := new(pb1.Req)
  fanout_1Req.Ids = res.Ids
	fanout_1Res, err := fanout_1Client.GetIds(context.Background(), fanout_1Req)
	if err == nil {
		res.Ids = fanout_1Res.Ids
	} else {
		log.Fatalf("Failed to retrieve ids from fanout_1: %v", err)
	}
  fanout_2Req := new(pb2.Req)
  fanout_2Req.Ids = res.Ids
	fanout_2Res, err := fanout_2Client.GetIds(context.Background(), fanout_2Req)
	if err == nil {
		res.Ids = fanout_2Res.Ids
	} else {
		log.Fatalf("Failed to retrieve ids from fanout_2: %v", err)
	}
  fanout_3Req := new(pb3.Req)
  fanout_3Req.Ids = res.Ids
	fanout_3Res, err := fanout_3Client.GetIds(context.Background(), fanout_3Req)
	if err == nil {
		res.Ids = fanout_3Res.Ids
	} else {
		log.Fatalf("Failed to retrieve ids from fanout_3: %v", err)
	}

	return res, err
}

func (s fanout_0Server) Check(ctx context.Context, req *healthpb.HealthCheckRequest) (*healthpb.HealthCheckResponse, error) {
	return &healthpb.HealthCheckResponse{Status: healthpb.HealthCheckResponse_SERVING}, nil
}

func (s fanout_0Server) Watch(req *healthpb.HealthCheckRequest, ws healthpb.Health_WatchServer) error {
	return status.Errorf(codes.Unimplemented, "health check via Watch not implemented")
}

