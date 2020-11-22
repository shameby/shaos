package main

import (
	"net"
	"log"
	_ "shaos/util/redisC"
	pb "shaos/proto/data"
	"shaos/data/heartbeat"
	"shaos/data/server"
	"shaos/data/conf"

	"google.golang.org/grpc"
)

func main() {
	go heartbeat.StartHearBeat()
	lis, err := net.Listen("tcp", ":"+*conf.BaseConfig.Port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	pb.RegisterDataServer(grpcServer, &server.DataServer{})
	grpcServer.Serve(lis)
}
