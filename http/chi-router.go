package router

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
)

var (
	chiNewRouter = chi.NewRouter()
)

type chiRouter struct{}

func NewChiRouter() Router {
	return &chiRouter{}
}

func (*chiRouter) GET(uri string, f func(w http.ResponseWriter, r *http.Request)) {

	chiNewRouter.Get(uri, f)
}
func (*chiRouter) POST(uri string, f func(w http.ResponseWriter, r *http.Request)) {
	chiNewRouter.Post(uri, f)
}
func (*chiRouter) SERVE(port string) {
	fmt.Printf("Chi HTTP server running on port : %v", port)
	http.ListenAndServe(port, chiNewRouter)
}
