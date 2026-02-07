package main

import (
	"flag"
	"fmt"
	"go/goProject/internal/app"
	"go/goProject/internal/routes"
	"net/http"
	"time"
)

func main() {

	var port int
	flag.IntVar(&port, "port", 8080, "Port to run the server on")
	flag.Parse()
	app, err := app.NewApplication()
	if err != nil {
		panic(err)
	}
	defer app.DB.Close()
	app.Logger.Println("Application running on port", port)

	r := routes.SetupRoutes(app)
	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", port),
		Handler:      r,
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	err = server.ListenAndServe()
	if err != nil {
		app.Logger.Fatalf("Server failed to start: %v", err)
	}
}
