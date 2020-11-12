package objectStream

import (
	"io"
	"context"

	pb "shaos/gateway/proto"

	"google.golang.org/grpc"
)

type PutStream struct {
	writer *io.PipeWriter
	c      chan error
	resp   *pb.PutReply
}

func NewPutStream(server, objName string) *PutStream {
	reader, writer := io.Pipe()
	c := make(chan error)
	ps := &PutStream{writer, c, nil}
	go func() {
		conn, err := grpc.Dial(server, grpc.WithInsecure())
		if err != nil {
			c <- err
			return
		}
		defer conn.Close()
		client := pb.NewDataClient(conn)
		buffer := make([]byte, 1048576) // 1M to send each time
		stream, err := client.Put(context.Background())
		stream.SendMsg(&pb.PutRequest{Name: objName})
		for {
			bytesReads, err := reader.Read(buffer)
			if err != nil {
				if err == io.EOF {
					if err := stream.Send(&pb.PutRequest{Data: buffer[:bytesReads]}); err != nil {
						c <- err
						return
					}
					break
				}
				c <- err
				return
			}
			if err := stream.Send(&pb.PutRequest{Data: buffer[:bytesReads]}); err != nil {
				c <- err
				return
			}
		}
		ps.resp, err = stream.CloseAndRecv()
		if err != nil {
			c <- err
			return
		}
		c <- err
	}()
	return ps
}

func (w *PutStream) Write(p []byte) (n int, err error) {
	return w.writer.Write(p)
}

func (w *PutStream) Close() error {
	w.writer.Close()
	return <-w.c
}

func (w PutStream) GetResp() *pb.PutReply {
	return w.resp
}