package object

import (
	"net/http"
	"context"

	"google.golang.org/grpc"
	"github.com/gorilla/mux"

	"shaos/gateway/heartbeat"
	pb "shaos/proto/data"
	. "shaos/gateway/types"
	"shaos/gateway/code"
)

func GetHandler(resp http.ResponseWriter, req *http.Request) {
	_, addr := heartbeat.ChooseRandomDataServer()
	conn, err := grpc.Dial(addr, grpc.WithInsecure())
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

