package internal

import (
	"github.com/getclasslabs/user/internal/handler"
	"github.com/getclasslabs/user/tools"
	"github.com/gorilla/mux"
	"net/http"
)

type Server struct {
	Router *mux.Router
}

func NewServer() *Server {
	r := mux.NewRouter()
	s := Server{r}

	s.serve()

	return &s
}

func (s *Server) serve() {

	s.Router.Path("/heartbeat").HandlerFunc(tools.PreRequest(handler.Heartbeat)).Methods(http.MethodGet)
	s.Router.Path("/create").HandlerFunc(tools.PreRequest(handler.CreateUser)).Methods(http.MethodPost)

}
