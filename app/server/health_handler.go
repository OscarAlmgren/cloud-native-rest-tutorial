package app

import "net/http"

func writeHealthy(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("."))
}

func writeUnhealthy(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusInternalServerError)
	w.Write([]byte("."))
}

func HandleLive(w http.ResponseWriter, _ *http.Request) {
	writeHealthy(w)
}

func (server *Server) HandleReady(w http.ResponseWriter, r *http.Request) {
	if err := server.db.DB().Ping(); err != nil {
		server.Logger().Fatal().Err(err).Msg("Fatal HandleReady error")
		writeUnhealthy(w)
		return
	}

	writeHealthy(w)
}
