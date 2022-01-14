package router

import (
	app "cloud-native/app/server"

	"github.com/go-chi/chi"
)

func New() *chi.Mux {
	router := chi.NewRouter()

	router.MethodFunc("GET", "/", app.HandleIndex)

	return router
}
