package main

import (
	"awesomeProject/database"
	"awesomeProject/handler"
	"awesomeProject/middleware"
	"net/http"
)

func main() {
	database.ConnectDB()

	router := http.NewServeMux()
	router.HandleFunc("POST /user", handler.CreateUser)
	router.HandleFunc("POST /user/login", handler.Login)

	protectedRouter := http.NewServeMux()
	protectedRouter.HandleFunc("GET /user/{id}", handler.GetUser)
	protectedRouter.HandleFunc("GET /me", handler.GetMe)
	protectedRouter.HandleFunc("POST /note", handler.CreateNote)
	protectedRouter.HandleFunc("GET /note", handler.GetNote)

	router.Handle("/", middleware.AuthenticateMiddleware(protectedRouter))
	server := http.Server{
		Addr:    ":8081",
		Handler: middleware.Logging(router),
	}
	err := server.ListenAndServe()
	if err != nil {
		return
	}
}
