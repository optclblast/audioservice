package api

import (
	"net/http"
	"time"

	"github.com/optclblast/filetagger/auth"
	"github.com/optclblast/filetagger/logger"

	"github.com/gorilla/mux"
)

type Server struct {
	*mux.Router
}

func NewServer() *Server {
	server := &Server{
		Router: mux.NewRouter(),
	}
	server.routes()
	return server
}

func (server *Server) routes() {
	server.HandleFunc("/startcommunication/public_rsakey", server.authInit()).Methods("GET")
}

func (server *Server) authInit() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(auth.HandleGetRSAKey()))
		logger.Logger(logger.LogEntry{
			DateTime: time.Now(),
			Level:    logger.DEBUG,
			Location: "ROUTS/auth_init() SCOPE",
			Content:  "Request at /startcommunication/public_rsakey",
		})
	}
}
