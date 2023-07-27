package main

import "fmt"

func main() {
	libraryServer := newLibraryServer()
	err := libraryServer.server.Run("localhost:8080")
	if err != nil {
		panic(fmt.Errorf("error occurred running library server: %w", err))
	}
}
