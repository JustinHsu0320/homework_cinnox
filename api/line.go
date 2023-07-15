package api

import (
	"fmt"
	"log"
	"net/http"

	db "homework_cinnox/db/mongo"

	"github.com/gin-gonic/gin"
	"github.com/line/line-bot-sdk-go/v7/linebot"
)

// Handle LINE message requests
func (server *Server) InsertMessage(ctx *gin.Context) {
	events, err := server.bot.ParseRequest(ctx.Request)
	if err != nil {
		if err == linebot.ErrInvalidSignature {
			ctx.Status(http.StatusBadRequest)
		} else {
			ctx.Status(http.StatusInternalServerError)
		}
		return
	}

	for _, event := range events {
		if event.Type == linebot.EventTypeMessage {
			// Handle incoming messages
			switch message := event.Message.(type) {
			case *linebot.TextMessage:
				// Debug
				// fmt.Printf("event: %#v\n", event)
				fmt.Printf("Source: %#v\n", event.Source)
				replyText := "You said: " + message.Text
				_, err := server.bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(replyText)).Do()
				if err != nil {
					log.Println(err)
				}

				// Save in DB
				arg := db.Message{
					UserID:  event.Source.UserID,
					Message: message.Text,
				}
				err = server.store.InsertMessage(arg)
				if err != nil {
					log.Println(err)
					return
				}

			}
		}
	}
}

// Push LINE message as response
func (server *Server) PushMessage(ctx *gin.Context) {
	// Parse request body to get the message content
	var request struct {
		Message string `json:"message"`
	}
	if err := ctx.BindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	// Create a new TextMessage to send
	message := linebot.NewTextMessage(request.Message)

	// Replace with your desired recipient ID or group ID
	recipientID := "U6e00d25d0266efcc9cdf86c688d93af4"

	// Send the message to the recipient
	if _, err := server.bot.PushMessage(recipientID, message).Do(); err != nil {
		log.Println(err)
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Message sent successfully"})
}
