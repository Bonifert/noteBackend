package router

import (
	"awesomeProject/handler"
	"awesomeProject/middleware"
	"net/http"
)

func SetUpRoutes() *http.ServeMux {
	r := http.NewServeMux()
	publicRoutes(r)
	protectedRoutes(r)
	return r
}

func publicRoutes(router *http.ServeMux) {
	router.HandleFunc("POST /user", handler.CreateUser)
	router.HandleFunc("POST /auth/login", handler.Login)
}

func protectedRoutes(mux *http.ServeMux) {
	protectedRouter := http.NewServeMux()

	// user related protected routes
	protectedRouter.HandleFunc("GET /user/me", handler.GetMe)
	protectedRouter.HandleFunc("GET /user/{id}", handler.GetUser)
	protectedRouter.HandleFunc("DELETE /user/me", handler.DeleteMe)
	protectedRouter.HandleFunc("PATCH /user/edit/username", handler.EditUsername)
	protectedRouter.HandleFunc("PATCH /user/edit/password", handler.EditPassword)

	// note related protected routes
	protectedRouter.HandleFunc("POST /note", handler.CreateNote)
	protectedRouter.HandleFunc("GET /note/{id}", handler.GetNote)
	protectedRouter.HandleFunc("DELETE /note/{id}", handler.DeleteNote)
	protectedRouter.HandleFunc("PATCH /note/content/{id}", handler.EditNoteContent)
	protectedRouter.HandleFunc("PATCH /note/title/{id}", handler.EditNoteTitle)

	mux.Handle("/", middleware.AuthenticateMiddleware(protectedRouter))
}
