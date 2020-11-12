package main

import (
	"net"
	"log"
	_ "shaos/util/redisC"
	pb "shaos/data/proto"
	"shaos/data/heartbeat"
	"shaos/data/server"

	"google.golang.org/grpc"
)

func main() {
	go heartbeat.StartHearBeat()
	lis, err := net.Listen("tcp", ":8081")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	pb.RegisterDataServer(grpcServer, &server.DataServer{})
	grpcServer.Serve(lis)
}
