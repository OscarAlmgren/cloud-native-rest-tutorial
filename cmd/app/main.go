package main

import (
	"fmt"
	"net/http"

	dbConn "cloud-native/adapter/gorm"
	"cloud-native/app/router"
	app "cloud-native/app/server"
	"cloud-native/config"
	"cloud-native/util/logger"
)

func main() {
	appConfig := config.AppConf()

	logger := logger.New(appConfig.Server.Debug)
	logger.Debug().Msgf("appConfig.Server.Debug: %t", appConfig.Server.Debug)

	db, err := dbConn.New(appConfig)
	if err != nil {
		logger.Fatal().Err(err).Msg("Database connection failed")
		return
	}
	if appConfig.Debug {
		db.LogMode(true)
	}

	application := app.New(logger, db)
	appRouter := router.New(application)

	address := fmt.Sprintf(":%d", appConfig.Server.Port)

	logger.Info().Msgf("Starting server %s", address)
	// log.Printf("Starting server %s\n", address)

	s := &http.Server{
		Addr:         address,
		Handler:      appRouter,
		ReadTimeout:  appConfig.Server.Timeoutread,
		WriteTimeout: appConfig.Server.Timeoutwrite,
		IdleTimeout:  appConfig.Server.Timeoutidle,
	}

	if err := s.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		logger.Fatal().Err(err).Msg("Server failed startup")
		// log.Fatal("Server startup failed")
	}

}

func Greet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello rest-api")
}
