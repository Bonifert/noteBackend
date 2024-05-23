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
	router.HandleFunc("POST /auth/login", handler.Login)

	protectedRouter := http.NewServeMux()
	protectedRouter.HandleFunc("GET /user/me", handler.GetMe)
	protectedRouter.HandleFunc("GET /user/{id}", handler.GetUser)
	protectedRouter.HandleFunc("POST /note", handler.CreateNote)
	protectedRouter.HandleFunc("GET /note/{id}", handler.GetNote)
	protectedRouter.HandleFunc("DELETE /user/me", handler.DeleteMe)

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
