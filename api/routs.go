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
	server.HandleFunc("/startcommunication/public_pgp", server.authInit()).Methods("GET")
	server.HandleFunc("/user/create_new_account", server.registerNewUser()).Methods("POST")
}

func (server *Server) authInit() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(auth.HandleGetRSAKey()))
		logger.Logger(logger.LogEntry{
			DateTime: time.Now(),
			Level:    logger.DEBUG,
			Location: "ROUTS/auth_init() SCOPE",
			Content:  "Request at /startcommunication/public_pgp",
		})
	}
}

func (server *Server) registerNewUser() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := r.URL.Query()
		userPublicKey := params.Get("pk")
		if userPublicKey == "" {
			logger.Logger(logger.LogEntry{
				DateTime: time.Now(),
				Level:    logger.INFO,
				Location: "ROUTS/registerNewUser() SCOPE",
				Content:  "An attempt to create new account. Request without RSA key",
			})
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("You cannot create a new account without providing your public RSA key!"))
			return
		}
		//DEBUG
		if userPublicKey == "DEBUG_KEY" {
			logger.Logger(logger.LogEntry{
				DateTime: time.Now(),
				Level:    logger.DEBUG,
				Location: "ROUTS/registerNewUser() SCOPE",
				Content:  "DEBUG REQUEST HANDLING (/user/create_new_account)",
			})
		}
		//DEBUG
		/*
			user := auth.UserAccount{
				Id:         db.GetLastUser() + 1,
				Login:      "1",
				Password:   "asdasd",
				Created_at: time.Now(),
				PublicKey:  "sdfsdf",
			}
			err := db.CreateNewUser(user)
			if err != nil {
				logger.Logger(logger.LogEntry{
					DateTime: time.Now(),
				Level:    logger.ERROR,
				Location: "ROUTS/auth_init() SCOPE",
					Content:  fmt.Sprintf("Failed to connect to the database: %s", err),
				})
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
		*/
	}
}
