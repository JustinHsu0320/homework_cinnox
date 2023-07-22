package main

import (
	"context"
	"log"

	"github.com/line/line-bot-sdk-go/v7/linebot"
	"github.com/spf13/cobra"
	"homework_cinnox/api"
	db "homework_cinnox/db/mongo"
	"homework_cinnox/util"
)

var rootCmd = &cobra.Command{Use: "linebot"}

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

	// Line Bot
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

func init() {
	// Define the 'TestCmd' command with the 'secret', 'token' flag
	var TestCmd = &cobra.Command{
		Use:   "testlinebot [secret] [token]",
		Short: "Test Line's official account",
		Run:   testFunc,
	}

	// Add the flags to the 'TestCmd' command directly
	TestCmd.Flags().StringP("secret", "s", "", "Secret for Line Bot (default: empty string)")
	TestCmd.Flags().StringP("token", "t", "", "Token for Line Bot (default: empty string)")

	// Add the Line Bot 'secret', 'token' command to the root command
	rootCmd.AddCommand(TestCmd)
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}

func testFunc(cmd *cobra.Command, args []string) {
	secret, err := cmd.Flags().GetString("secret")
	if err != nil {
		log.Fatal(err)
		return
	}
	token, err := cmd.Flags().GetString("token")
	if err != nil {
		log.Fatal(err)
		return
	}

	log.Printf("secret = %s, token = %s", secret, token)

	// LINE Bot
	bot, err := linebot.New(secret, token)
	log.Printf("bot: %#v", bot)
	if err != nil {
		log.Fatal(err)
		return
	}

	log.Println("OK")
}
