package main

import (
	"awesomeProject/pkg/database"
	"awesomeProject/pkg/middleware"
	"awesomeProject/pkg/router"
	"net/http"
)

func main() {
	database.ConnectDB()
	routes := router.SetUpRoutes()

	server := http.Server{
		Addr:    ":8081",
		Handler: middleware.Logging(routes),
	}
	err := server.ListenAndServe()
	if err != nil {
		return
	}
}
