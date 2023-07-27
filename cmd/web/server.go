package main

import (
	"fmt"
	"github.com/BieLuk/library-backend/src/config"
	"github.com/BieLuk/library-backend/src/controller"
	"github.com/BieLuk/library-backend/src/db"
	"github.com/gin-gonic/gin"
)

type libraryServer struct {
	server *gin.Engine

	bookController controller.BookController
}

func runLibraryServer() {
	appConfig, err := config.LoadConfig(".")
	if err != nil {
		panic(fmt.Errorf("error occurred reading config file"))
	}
	if err := db.Init(appConfig); err != nil {
		panic(fmt.Errorf("error initializing database: %w", err))
	}

	libServer := &libraryServer{
		server:         gin.Default(),
		bookController: controller.NewBookController(),
	}

	libServer.routes()

	if err := libServer.server.Run(fmt.Sprintf(":%s", appConfig.ServerPort)); err != nil {
		panic(fmt.Errorf("error occurred running library server: %w", err))
	}
}
