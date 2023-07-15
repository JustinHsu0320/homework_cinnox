package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	db "homework_cinnox/db/mongo"
)

type createUserRequest struct {
	ID       string `json:"id"`
	Username string `json:"username" binding:"required,alphanum"`
	Email    string `json:"email" binding:"required,email"`
}

type userResponse struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

func newUserResponse(user db.User) userResponse {
	return userResponse{
		ID:       user.ID,
		Username: user.Username,
		Email:    user.Email,
	}
}

func (server *Server) InsertUser(ctx *gin.Context) {
	var req createUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.User{
		ID:       req.ID,
		Username: req.Username,
		Email:    req.Email,
	}

	err := server.store.InsertUser(arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	rsp := newUserResponse(arg)
	ctx.JSON(http.StatusOK, rsp)
}
