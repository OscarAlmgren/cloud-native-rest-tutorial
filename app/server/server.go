package server

import (
	"cloud-native/util/logger"

	"github.com/jinzhu/gorm"
)

type Server struct {
	logger *logger.Logger
	db     *gorm.DB
}

func New(logger *logger.Logger, db *gorm.DB) *Server {
	return &Server{
		logger: logger,
		db:     db,
	}
}

func (server *Server) Logger() *logger.Logger {
	return server.logger
}
