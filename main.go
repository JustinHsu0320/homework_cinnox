package main

import (
	"log"

	"github.com/JustinHsu0320/homework_cinnox/api"
	db "github.com/JustinHsu0320/homework_cinnox/db/mongo"
	"github.com/JustinHsu0320/homework_cinnox/util"
)

func main() {
	// env
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Println("cannot load config")
	}

	// mongodb conn pool
	pool, err := db.GetMongoDBClientPool(config.DBSource)
	if err != nil {
		log.Println("cannot connect mongodb")
	}

	store := db.NewStore(pool)

	runGinServer(config, store)
}

func runGinServer(config util.Config, store db.Store) {
	server, err := api.NewServer(config, store)
	if err != nil {
		log.Println("cannot create server")
	}

	err = server.Start(config.HTTPServerAddress)
	if err != nil {
		log.Println("cannot start server")
	}
}
