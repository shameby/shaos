package main

import (
	"log"
	"net/http"

	"shaos/gateway/types"
	"shaos/gateway/router"
	_ "shaos/util/redisC"
	"shaos/gateway/conf"

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
	r := router.GenRouter()
	r.Use(TokenMiddleware)
	log.Println(*conf.BaseConfig.ServerName + " start listening on " + *conf.BaseConfig.Port)
	server := &http.Server{Addr: ":" + *conf.BaseConfig.Port, Handler: r}
	defer server.Close()
	log.Fatal(server.ListenAndServe())
}
