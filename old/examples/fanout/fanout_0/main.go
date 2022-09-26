package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "github.com/hanapedia/microservice-topologies/fanout/fanout_0/pb_fanout_0"
	pb1 "github.com/hanapedia/microservice-topologies/fanout/fanout_0/pb_fanout_1"
	pb2 "github.com/hanapedia/microservice-topologies/fanout/fanout_0/pb_fanout_2"
	pb3 "github.com/hanapedia/microservice-topologies/fanout/fanout_0/pb_fanout_3"
)

type fanout_0Server struct {
	pb.UnimplementedFanout_0Server
}

const (
	ListenPort           = "3001"
	FanoutClient1Address = "localhost:3002"
	FanoutClient2Address = "localhost:3003"
	FanoutClient3Address = "localhost:3004"
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

  router := gin.Default()
  router.GET("/", handler)
  router.Run(fmt.Sprintf(":%s", listenPort))
}

func handler(c *gin.Context) {
  var res []int32

  fanout_1Req := new(pb1.Req)
  fanout_1Req.Ids = res
	fanout_1Res, err := fanout_1Client.GetIds(context.Background(), fanout_1Req)
	if err == nil {
		res = fanout_1Res.Ids
	} else {
		log.Fatalf("Failed to retrieve ids from fanout_1: %v", err)
    c.JSON(http.StatusBadRequest, gin.H{
      "message": err,
    })
	}
  fanout_2Req := new(pb2.Req)
  fanout_2Req.Ids = res
	fanout_2Res, err := fanout_2Client.GetIds(context.Background(), fanout_2Req)
	if err == nil {
		res = fanout_2Res.Ids
	} else {
		log.Fatalf("Failed to retrieve ids from fanout_2: %v", err)
    c.JSON(http.StatusBadRequest, gin.H{
      "message": err,
    })
	}
  fanout_3Req := new(pb3.Req)
  fanout_3Req.Ids = res
	fanout_3Res, err := fanout_3Client.GetIds(context.Background(), fanout_3Req)
	if err == nil {
		res = fanout_3Res.Ids
	} else {
		log.Fatalf("Failed to retrieve ids from fanout_3: %v", err)
    c.JSON(http.StatusBadRequest, gin.H{
      "message": err,
    })
	}
  c.JSON(http.StatusBadRequest, gin.H{
    "message": res,
  })
}
