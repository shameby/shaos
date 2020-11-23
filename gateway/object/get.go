package object

import (
	"net/http"
	"context"

	"github.com/gorilla/mux"

	"shaos/gateway/heartbeat"
	pb "shaos/proto/data"
	. "shaos/gateway/types"
	"shaos/gateway/code"
	"shaos/util/grpcP"
)

func GetHandler(resp http.ResponseWriter, req *http.Request) {
	_, addr := heartbeat.ChooseRandomDataServer()
	conn, err := grpcP.GetConn(addr)
	if err != nil {
		AjaxResp(resp, code.ModulesFuncErr, err.Error(), "")
		return
	}
	defer conn.Close()
	client := pb.NewDataClient(conn)
	vars := mux.Vars(req)
	res, err := client.Get(context.Background(), &pb.GetRequest{Name: vars["name"]})
	if err != nil {
		AjaxResp(resp, code.ModulesFuncErr, err.Error(), "")
		return
	}
	AjaxResp(resp, http.StatusOK, "success", res.Data)
	return
}

