package commands

import (
	"context"
	// "time"
	// "github.com/Hello-Storage/hello-back/internal/entity"

	"github.com/Hello-Storage/hello-storage-proxy/internal/config"
	"github.com/Hello-Storage/hello-storage-proxy/internal/event"
	"github.com/Hello-Storage/hello-storage-proxy/internal/server"
)

var log = event.Log

func Start() {
	// init logger
	config.InitLogger()

	// load env
	err := config.LoadEnv()
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	// connect db and define enum types
	err = config.ConnectDB()
	if err != nil {
		log.Fatal("cannot connect to DB and create enums:", err)
	}

	config.InitDb()

	// Pass this context down the chain.
	cctx, cancel := context.WithCancel(context.Background())

	server.Start(cctx)

	// Cancel the context when the server stops
	cancel()
}
