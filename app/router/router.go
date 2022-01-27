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

	return chiRouter
}
