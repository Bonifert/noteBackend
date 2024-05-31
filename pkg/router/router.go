package router

import (
	handler2 "awesomeProject/pkg/handler"
	"awesomeProject/pkg/middleware"
	"net/http"
)

func SetUpRoutes() *http.ServeMux {
	r := http.NewServeMux()
	publicRoutes(r)
	protectedRoutes(r)
	return r
}

func publicRoutes(router *http.ServeMux) {
	router.HandleFunc("POST /user", handler2.CreateUser)
	router.HandleFunc("POST /auth/login", handler2.Login)
}

func protectedRoutes(mux *http.ServeMux) {
	protectedRouter := http.NewServeMux()

	// user related protected routes
	protectedRouter.HandleFunc("GET /user/me", handler2.GetMe)
	protectedRouter.HandleFunc("GET /user/{id}", handler2.GetUser)
	protectedRouter.HandleFunc("DELETE /user/me", handler2.DeleteMe)
	protectedRouter.HandleFunc("PATCH /user/edit/username", handler2.EditUsername)
	protectedRouter.HandleFunc("PATCH /user/edit/password", handler2.EditPassword)

	// note related protected routes
	protectedRouter.HandleFunc("POST /note", handler2.CreateNote)
	protectedRouter.HandleFunc("GET /note/{id}", handler2.GetNote)
	protectedRouter.HandleFunc("DELETE /note/{id}", handler2.DeleteNote)
	protectedRouter.HandleFunc("PATCH /note/content/{id}", handler2.EditNoteContent)
	protectedRouter.HandleFunc("PATCH /note/title/{id}", handler2.EditNoteTitle)

	mux.Handle("/", middleware.AuthenticateMiddleware(protectedRouter))
}
