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

	pb "github.com/hanapedia/microservice-topologies/fanout/test_templates/fanout_0/pb_fanout_0"
	pbx "github.com/hanapedia/microservice-topologies/fanout/test_templates/fanout_0/pb_fanout_x"
)

type fanout_0Server struct {
	pb.UnimplementedFanout_0Server
}

const (
	ListenPort           = "3001"
	FanoutClientXAddress = "localhost:3002"
)

var fanoutXClient pbx.FanoutXClient
var conn *grpc.ClientConn

func main() {
	listenPort := ListenPort
	if os.Getenv("LISTEN_PORT") != "" {
		listenPort = os.Getenv("LISTEN_PORT")
	}
	fanoutClientXAddress := FanoutClientXAddress
	if os.Getenv("FANOUT_CLIENT_X_ADDRESS") != "" {
		fanoutClientXAddress = os.Getenv("FANOUT_CLIENT_X_ADDRESS")
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

	conn, err = grpc.Dial(fanoutClientXAddress, optsClient...)
	if err != nil {
		log.Fatalf("Cannot establish connection with the server: %v", err)
	}
	log.Printf("Dialed and established connection with %v", fanoutClientXAddress)
	defer conn.Close()

	fanoutXClient = pbx.NewFanoutXClient(conn)

	log.Printf("Starting server at port %v", listenPort)
	grpcServer.Serve(lis)
}

func (s fanout_0Server) GetIds(ctx context.Context, req *pb.Req) (*pb.Res, error) {
	res := new(pb.Res)
	res.Ids = req.Ids

  fanoutXReq := new(pbx.Req)
  fanoutXReq.Ids = res.Ids
	fanoutXRes, err := fanoutXClient.GetIds(context.Background(), fanoutXReq)
	if err == nil {
		res.Ids = fanoutXRes.Ids
	} else {
		log.Fatalf("Failed to retrieve ids: %v", err)
	}

	return res, err
}

func (s fanout_0Server) Check(ctx context.Context, req *healthpb.HealthCheckRequest) (*healthpb.HealthCheckResponse, error) {
	return &healthpb.HealthCheckResponse{Status: healthpb.HealthCheckResponse_SERVING}, nil
}

func (s fanout_0Server) Watch(req *healthpb.HealthCheckRequest, ws healthpb.Health_WatchServer) error {
	return status.Errorf(codes.Unimplemented, "health check via Watch not implemented")
}
