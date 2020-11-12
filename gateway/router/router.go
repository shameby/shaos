package router

import (
	"github.com/gorilla/mux"
)

type group struct {
	name    string
	prefix  string
	handler func(*mux.Router)
}

type groupArr []group

var groups = groupArr{
	group{name: "objects", prefix: "/objects", handler: ObjectRoutes},
}

// GenRouter
func GenRouter() *mux.Router {
	r := mux.NewRouter()
	for _, group := range groups {
		s := r.PathPrefix(group.prefix).Subrouter()
		group.handler(s)
	}
	return r
}
