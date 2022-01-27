package router

import (
	"cloud-native/app/requestlog"
	middleware "cloud-native/app/router/middleware"
	server "cloud-native/app/server"

	"github.com/go-chi/chi"
)

func New(server *server.Server) *chi.Mux {
	logger := server.Logger()
	chiRouter := chi.NewRouter()

	chiRouter.Method("GET", "/", requestlog.NewHandler(server.HandleIndex, logger))
	chiRouter.Get("/healthz/readiness", server.HandleLive)

	chiRouter.Route("/api/v1", func(r chi.Router) {
		r.Use(middleware.ContentTypeJson)
		r.Method("GET", "/books", requestlog.NewHandler(server.HandleListBooks, logger))
		r.Method("POST", "/books", requestlog.NewHandler(server.HandleCreateBook, logger))
		r.Method("GET", "/books/{id}", requestlog.NewHandler(server.HandleReadBook, logger))
		r.Method("PUT", "/books/{id}", requestlog.NewHandler(server.HandleUpdateBook, logger))
		r.Method("DELETE", "/books/{id}", requestlog.NewHandler(server.HandleDeleteBook, logger))
	})

	return chiRouter
}
