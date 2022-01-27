package router

import (
	"cloud-native/app/requestlog"
	server "cloud-native/app/server"

	"github.com/go-chi/chi"
)

func New(server *server.Server) *chi.Mux {
	logger := server.Logger()
	chiRouter := chi.NewRouter()

	chiRouter.Method("GET", "/", requestlog.NewHandler(server.HandleIndex, logger))
	chiRouter.Get("/healthz/readiness", server.HandleLive)

	chiRouter.Method("GET", "/books", requestlog.NewHandler(server.HandleListBooks, logger))
	chiRouter.Method("POST", "books", requestlog.NewHandler(server.HandleCreateBook, logger))
	chiRouter.Method("GET", "/books/{id}", requestlog.NewHandler(server.HandleReadBook, logger))
	chiRouter.Method("PUT", "/books/{id}", requestlog.NewHandler(server.HandleDeleteBook, logger))

	return chiRouter
}
