package main

import (
	"context"
	"log"
	"net/http"
  "os"
  "fmt"

	"github.com/gin-gonic/gin"

	pb "github.com/hanapedia/microservice-topologies/chain_example/root/pb_chain_1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
  ListenPort = "3000"
  ChainNextAddress = "localhost:3001"
)

var chainClient pb.Chain_1Client

func main() {
	listenPort := ListenPort
	if os.Getenv("LISTEN_PORT") != "" {
		listenPort = os.Getenv("LISTEN_PORT")
	}
  chainNextAddress := ChainNextAddress
	if os.Getenv("CHAIN_NEXT_ADDRESS") != "" {
		chainNextAddress = os.Getenv("CHAIN_NEXT_ADDRESS")
	}
  var opts []grpc.DialOption
  opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))

  conn, err := grpc.Dial(chainNextAddress, opts...)
  if err != nil {
    log.Printf("Cannot establish connection with the server: %v", err)
  }
  defer conn.Close()

  chainClient = pb.NewChain_1Client(conn)

  router := gin.Default()
  router.GET("/", handler)
  router.Run(fmt.Sprintf(":%s", listenPort))
}

func handler(c *gin.Context) {
  ids, err := chainClient.GetIds(context.Background(), &pb.Req{Ids: []int32{25}})
  if err != nil {
    c.JSON(http.StatusBadRequest, gin.H{
      "message": err,
    })
  }
  c.JSON(http.StatusBadRequest, gin.H{
    "message": ids,
  })
}
