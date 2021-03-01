package router

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type chiRouter struct {
}

var (
	muxDispatcher = mux.NewRouter()
)

func NewChiRouter() Router {
	return &chiRouter{}
}

func (*muxRouter) GET(uri string, f func(w http.ResponseWriter, r *http.Request)) {
	muxDispatcher.HandleFunc(uri, f).Methods("GET")
}

func (*muxRouter) POST(uri string, f func(w http.ResponseWriter, r *http.Request)) {
	muxDispatcher.HandleFunc(uri, f).Methods("POST")

}

func (*muxRouter) SERVE(port string) {
	fmt.Printf("Mux HTTP serveer running on port %v", port)
	http.ListenAndServe(port, muxDispatcher)
}
