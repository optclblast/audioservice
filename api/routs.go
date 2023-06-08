package api

import (
	"github.com/gin-gonic/gin"
	"github.com/optclblast/filetagger/db"
)

type Server struct {
	store  *db.Store
	router *gin.Engine
}

func NewServer(store *db.Store) *Server {
	server := &Server{store: store}
	router := gin.Default()

	router.GET("/authInit", server.authInit)
	router.POST("/createaccount", server.createAccount)

	server.router = router
	return server
}
