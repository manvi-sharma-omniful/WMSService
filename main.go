package main

import (
	appinit "awesomeProject/Project/WMS/init"
	"awesomeProject/Project/WMS/router"
	"context"
	"github.com/omniful/go_commons/config"
	"github.com/omniful/go_commons/http"
	"github.com/omniful/go_commons/log"
	"github.com/omniful/go_commons/shutdown"
	"os"
	"time"
)

func main() {
	err := os.Setenv("CONFIG_SOURCE", "local")
	if err != nil {
		log.Panicf("Error while initialising config, err: %v", err)
		panic(err)
	}
	// Initialize config
	err = config.Init(time.Second * 10)
	if err != nil {
		log.Panicf("Error while initialising config, err: %v", err)
		panic(err)
	}

	ctx, err := config.TODOContext()
	if err != nil {
		log.Panicf("Error while getting context from config, err: %v", err)
		panic(err)
	}

	appinit.Initialize(ctx)

	runHttpServer(ctx)
}

func runHttpServer(ctx context.Context) {
	server := http.InitializeServer(config.GetString(ctx, "server.port"), 10*time.Second, 10*time.Second, 70*time.Second)

	// Initialize routes
	err := router.InternalRoutes(ctx, server)
	if err != nil {
		log.Errorf(err.Error())
		panic(err)
	}
	log.Infof("Starting server on port" + config.GetString(ctx, "server.port"))

	err = server.StartServer("warehouse-service")
	if err != nil {
		log.Errorf(err.Error())
		panic(err)
	}

	<-shutdown.GetWaitChannel()
}
