package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (server *Server) GetUsers(ctx *gin.Context) {

	// Execute the aggregation query
	users, err := server.store.GetUsers()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	// Return the list of unique users
	ctx.JSON(http.StatusOK, gin.H{"users": users})
}
