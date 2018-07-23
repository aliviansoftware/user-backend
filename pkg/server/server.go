package server

import (
	"log"
	"net/http"
	"os"
	"user-backend/pkg"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

type Server struct {
	router *mux.Router
	config *pkg.ServerConfig
}

func NewServer(u pkg.UserService, config *pkg.Config) *Server {
	s := Server{
		router: mux.NewRouter(),
		config: config.Server}

	a := authHelper{config.Auth.Secret}
	NewUserRouter(u, s.getSubrouter("/user"), &a)
	return &s
}

func (s *Server) Start() {
	log.Println("Listening on port " + s.config.Port)
	if err := http.ListenAndServe(s.config.Port, handlers.LoggingHandler(os.Stdout, s.router)); err != nil {
		log.Fatal("http.ListenAndServe: ", err)
	}
}

func (s *Server) getSubrouter(path string) *mux.Router {
	return s.router.PathPrefix(path).Subrouter()
}
