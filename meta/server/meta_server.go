package server

import (
	"context"
	"errors"

	pb "shaos/proto/meta"
	"shaos/meta/types"
)

type MetaServer struct{}

func (ms *MetaServer) CreateBucket(ctx context.Context, in *pb.CreateBucketReq) (*pb.CreateBucketReply, error) {
	if err := createBucket(in.Name); err != nil {
		return nil, err
	}
	return &pb.CreateBucketReply{Name: in.Name}, nil
}

func (ms *MetaServer) GetMetaById(ctx context.Context, in *pb.GetByIdRequest) (*pb.GetByIdReply, error) {
	col := getBucket(in.Bucket)
	if col == nil {
		return nil, errors.New("empty bucket name")
	}
	data, err := queryById(in.Id, col)
	if err != nil {
		return nil, err
	}
	return &pb.GetByIdReply{
		Id:           data.Id,
		FileName:     data.FileOriginName,
		ServerAppKey: data.DataServerAppKey,
		Version:      data.Version,
	}, nil
}

func (ms *MetaServer) PutMeta(ctx context.Context, in *pb.PutMetaRequest) (*pb.PutMetaReply, error) {
	col := getBucket(in.Bucket)
	if col == nil {
		return nil, errors.New("empty bucket name")
	}
	meta := types.NewMetaData(in.FileName, in.ServerAppKey)
	err := saveById(col, meta)
	if err != nil {
		return nil, err
	}
	return &pb.PutMetaReply{Id: meta.Id}, nil
}
