package main

import (
	"net"
	"log"

	"shaos/meta/heartbeat"
	"shaos/meta/conf"
	"shaos/meta/server"
	pb "shaos/proto/meta"

	"google.golang.org/grpc"
)

func main() {
	go heartbeat.StartHearBeat()
	lis, err := net.Listen("tcp", ":"+*conf.BaseConfig.Port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	pb.RegisterMetaServer(grpcServer, &server.MetaServer{})
	grpcServer.Serve(lis)
}