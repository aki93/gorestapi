package router

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// variable para utilizar la dependencia mux
var muxNewRouter = mux.NewRouter()

// implementacion de la interface
type muxRouter struct{}

// Constructor para la instancia de Router (router.go)
// esta devuelve el struct que implementa la interfaz
func NewMuxRouter() Router {
	return &muxRouter{}
}

//implementacion de los metodos de router.go

func (*muxRouter) GET(uri string, f func(w http.ResponseWriter, r *http.Request)) {

	muxNewRouter.HandleFunc(uri, f).Methods("GET")
}
func (*muxRouter) POST(uri string, f func(w http.ResponseWriter, r *http.Request)) {

	muxNewRouter.HandleFunc(uri, f).Methods("POST")

}
func (*muxRouter) SERVE(port string) {
	fmt.Printf("Mux HTTP server running on port : %v", port)
	http.ListenAndServe(port, muxNewRouter)
}
