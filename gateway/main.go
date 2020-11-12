package main

import (
	"log"
	"net/http"

	"shaos/gateway/types"
	"shaos/gateway/router"
	"shaos/gateway/heartbeat"
	_ "shaos/util/redisC"

	"github.com/json-iterator/go"
)

func TokenMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")
		if token == "Token" {
			next.ServeHTTP(w, r)
		} else {
			var result = types.AjaxResult{
				Code:    400,
				Message: "Token auth fail",
			}
			data, _ := jsoniter.MarshalToString(result)
			http.Error(w, data, http.StatusBadRequest)
		}
	})
}

func main() {
	go heartbeat.ListenHeartbeat()

	r := router.GenRouter()
	r.Use(TokenMiddleware)
	log.Printf("start listening on " + "8080")
	server := &http.Server{Addr: ":" + "8080", Handler: r}
	defer server.Close()
	log.Fatal(server.ListenAndServe())
}
