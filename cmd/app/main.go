package main

import (
	"fmt"
	"log"
	"net/http"

	"cloud-native/config"
)

func main() {
	appConfig := config.AppConf()

	mux := http.NewServeMux()
	mux.HandleFunc("/", Greet)

	address := fmt.Sprintf(":%d", appConfig.Server.Port)

	log.Printf("Starting server %s\n", address)

	s := &http.Server{
		Addr:         address,
		Handler:      mux,
		ReadTimeout:  appConfig.Server.Timeoutread,
		WriteTimeout: appConfig.Server.Timeoutwrite,
		IdleTimeout:  appConfig.Server.Timeoutidle
	}

	if err := s.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatal("Server startup failed")
	}

}

func Greet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello rest-api")
}
