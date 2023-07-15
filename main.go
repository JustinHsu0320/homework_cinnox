package main

import (
	"context"
	"log"

	"homework_cinnox/api"
	db "homework_cinnox/db/mongo"
	"homework_cinnox/util"

	"github.com/line/line-bot-sdk-go/v7/linebot"
)

func main() {

	ctx := context.Background()

	// env
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config")
	}

	// MongoDB Conn Pool
	pool, err := db.GetMongoDBClientPool(ctx, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect mongodb")
	}
	defer pool.Disconnect(ctx)

	// LINE Bot
	bot, err := linebot.New(config.LineChannelSecret, config.LineChannelAccessToken)
	if err != nil {
		log.Fatal(err)
	}

	store := db.NewStore(pool)

	runGinServer(config, store, bot)
}

func runGinServer(config util.Config, store *db.MongoStore, bot *linebot.Client) {
	server, err := api.NewServer(config, store, bot)
	if err != nil {
		log.Println("cannot create server")
	}

	err = server.Start(config.HTTPServerAddress)
	if err != nil {
		log.Println("cannot start server")
	}
}
