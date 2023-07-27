package main

import (
	"github.com/BieLuk/library-backend/src/controller"
	"github.com/gin-gonic/gin"
)

type LibraryServer interface {
}

type libraryServer struct {
	server *gin.Engine

	bookController controller.BookController
}

func newLibraryServer() *libraryServer {

	libServer := &libraryServer{
		server:         gin.Default(),
		bookController: controller.NewBookController(),
	}

	libServer.routes()

	return libServer
}
