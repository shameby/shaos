package server

import (
	"os"
	"io"
	"io/ioutil"
	"context"

	pb "shaos/proto/data"
)

type DataServer struct{}

func (ds *DataServer) Get(ctx context.Context, in *pb.GetRequest) (*pb.GetReply, error) {
	f, err := os.Open("./objects/" + in.Name)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	fd, err := ioutil.ReadAll(f)
	if err != nil {
		return nil, err
	}
	return &pb.GetReply{Data: fd}, nil
}

func (ds *DataServer) Put(dps pb.Data_PutServer) error {
	req := &pb.PutRequest{}
	dps.RecvMsg(req)
	f, e := os.Create("./objects/" + req.Name)
	if e != nil {
		return e
	}
	defer f.Close()
	for {
		//从流中获取消息
		stream, err := dps.Recv()
		if err == io.EOF {
			return dps.SendAndClose(&pb.PutReply{Name:req.Name})
		}
		if err != nil {
			return err
		}
		f.Write(stream.Data)
	}
}
