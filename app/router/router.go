package router

import (
	"cloud-native/app/requestlog"
	server "cloud-native/app/server"

	"github.com/go-chi/chi"
)

func New(server *server.Server) *chi.Mux {
	logger := server.Logger()
	router := chi.NewRouter()

	router.Method("GET", "/", requestlog.NewHandler(server.HandleIndex, logger))

	return router
}
