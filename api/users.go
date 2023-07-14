package api

import (
	"net/http"

	db "github.com/JustinHsu0320/homework_cinnox/db/mongo"
	"github.com/gin-gonic/gin"
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

	err := db.InsertUser(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	rsp := newUserResponse(arg)
	ctx.JSON(http.StatusOK, rsp)
}
