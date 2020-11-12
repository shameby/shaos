package object

import (
	"io"
	"net/http"
	"errors"

	"github.com/gorilla/mux"

	"shaos/gateway/heartbeat"
	"shaos/gateway/object_stream"
	"shaos/gateway/code"
	. "shaos/gateway/types"
	pb "shaos/gateway/proto"
)

func PutHandler(resp http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	grpcResp, e := storeObject(req.Body, vars["name"])
	if e != nil {
		AjaxResp(resp, code.ModulesFuncErr, e.Error(), grpcResp.Name)
	}
	AjaxResp(resp, http.StatusOK, "", grpcResp.Name)
	return
}

func storeObject(r io.Reader, objName string) (*pb.PutReply, error) {
	stream, err := putStream(objName)
	if err != nil {
		return nil, err
	}
	io.Copy(stream, r)
	err = stream.Close()
	if err != nil {
		return nil, err
	}
	return stream.GetResp(), nil
}

func putStream(objName string) (*objectStream.PutStream, error) {
	server := heartbeat.ChooseRandomDataServer()
	if server == "" {
		return nil, errors.New("cannot find any dataServer")
	}
	return objectStream.NewPutStream(server, objName), nil
}
