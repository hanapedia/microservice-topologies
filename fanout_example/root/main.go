package main

import (
	"context"
	"log"

	pb "github.com/hanapedia/microservice-topologies/fanout_example/root/pb_fanout_0"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
  var opts []grpc.DialOption
  opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))

  conn, err := grpc.Dial("localhost:4001", opts...)
  if err != nil {
    log.Printf("Cannot establish connection with the server: %v", err)
  }
  defer conn.Close()

  client := pb.NewFanout_0Client(conn)
  ids, err := client.GetIds(context.Background(), &pb.Req{Ids: []int32{25}})
  if err != nil {
    log.Printf("failed to retrieve ids: %v", err)
  }
  log.Printf("ids: %v", ids)
}
