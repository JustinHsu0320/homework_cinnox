package api

import (
	db "homework_cinnox/db/mongo"
	"homework_cinnox/util"

	"github.com/gin-gonic/gin"
	"github.com/line/line-bot-sdk-go/v7/linebot"
)

// Server serves HTTP requests for our service.
type Server struct {
	config util.Config
	store  *db.MongoStore
	bot    *linebot.Client
	router *gin.Engine
}

// NewServer creates a new HTTP server and set up routing.
func NewServer(config util.Config, store *db.MongoStore, bot *linebot.Client) (*Server, error) {
	server := &Server{
		config: config,
		store:  store,
		bot:    bot,
	}

	server.setupRouter()
	return server, nil
}

func (server *Server) setupRouter() {
	router := gin.Default()

	router.POST("/users", server.InsertUser)
	router.POST("/webhook", server.InsertMessage)
	router.POST("/messages/send", server.PushMessage)
	router.GET("/messages/users", server.GetUsers)

	server.router = router
}

// Start runs the HTTP server on a specific address.
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
