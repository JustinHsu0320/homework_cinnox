package api

import (
	"fmt"

	db "github.com/JustinHsu0320/homework_cinnox/db/mongo/store.go"
	"github.com/JustinHsu0320/homework_cinnox/util"
	"github.com/gin-gonic/gin"
)

// Server serves HTTP requests for our service.
type Server struct {
	config util.Config
	store  db.Store
	router *gin.Engine
}

// NewServer creates a new HTTP server and set up routing.
func NewServer(config util.Config, store db.Store) (*Server, error) {
	server := &Server{
		config: config,
		store:  store,
	}

	server.setupRouter()
	return server, nil
}

func (server *Server) setupRouter() {
	router := gin.Default()

	router.POST("/users", server.InsertUser)
	// router.GET("/users/:username", server.FindUser)
	// router.PUT("/users/login", server.UpdateUserEmail)
	// router.DELETE("/users/:username", server.DeleteUser)

	server.router = router
}

// Start runs the HTTP server on a specific address.
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
