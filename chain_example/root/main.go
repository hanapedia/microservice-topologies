package main

import (
	"context"
	"log"

	pb "github.com/hanapedia/microservice-topologies/chain_example/root/pb_chain_1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
  var opts []grpc.DialOption
  opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))

  conn, err := grpc.Dial("localhost:3001", opts...)
  if err != nil {
    log.Printf("Cannot establish connection with the server: %v", err)
  }
  defer conn.Close()

  client := pb.NewChain_1Client(conn)
  ids, err := client.GetIds(context.Background(), &pb.Req{Ids: []int32{25}})
  if err != nil {
    log.Printf("failed to retrieve ids: %v", err)
  }
  log.Printf("ids: %v", ids)
}
