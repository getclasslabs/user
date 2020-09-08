package internal

import (
	"github.com/getclasslabs/go-tools/pkg/request"
	"github.com/getclasslabs/user/internal/handler"
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

	s.Router.Path("/heartbeat").HandlerFunc(request.PreRequest(handler.Heartbeat)).Methods(http.MethodGet)
	s.Router.Path("/create").HandlerFunc(request.PreRequest(handler.CreateUser)).Methods(http.MethodPost)
	s.Router.Path("/profile").HandlerFunc(request.PreRequest(handler.CreateProfile)).Methods(http.MethodPut)

	s.Router.Path("/edit").HandlerFunc(request.PreRequest(handler.EditProfile)).Methods(http.MethodPut)

	s.Router.Path("/login").HandlerFunc(request.PreRequest(handler.Login)).Methods(http.MethodPost)

	s.Router.Path("/u/{nickname}").HandlerFunc(request.PreRequest(handler.GetUser)).Methods(http.MethodPost)

	s.Router.Path("/search/teacher").HandlerFunc(request.PreRequest(handler.SearchTeacher)).Methods(http.MethodGet)
}
