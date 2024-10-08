package main

import (
	"net/http"

	"os"
	"os/signal"
	"pingcrm/pkg/server"

	"log"
)

func main() {
	c := server.NewContainer()
	defer func() {
		log.Print("Shutting down server")
		c.Shutdown()
	}()

	// Start the server
	go func() {
		routes := server.SetupRoutes(c)
		http.ListenAndServe("localhost:3000", routes)
	}()

	// Wait for interrupt signal to gracefully shut down the server with a timeout of 10 seconds.
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	signal.Notify(quit, os.Kill)
	<-quit
}
