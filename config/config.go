package config

import (
	"time"

	"cloud-native/util/logger"

	"github.com/joeshaw/envdecode"
)

type Conf struct {
	Debug  bool `env:"DEBUG,required"`
	Server serverConf
	Db     dbConf
}

type serverConf struct {
	Port         int           `env:"SERVER_PORT,required"`
	Timeoutread  time.Duration `env:"SERVER_TIMEOUT_READ,required"`
	Timeoutwrite time.Duration `env:"SERVER_TIMEOUT_WRITE,required"`
	Timeoutidle  time.Duration `env:"SERVER_TIMEOUT_IDLE,required"`
	Debug        bool          `env:"DEBUG,required"`
}

type dbConf struct {
	Host     string `env:"DB_HOST,required"`
	Port     int    `env:"DB_PORT,required"`
	Username string `env:"DB_USER,required"`
	Password string `env:"DB_PASS,required"`
	DbName   string `env:"DB_NAME,required"`
}

func ServerConf() *Conf {
	var c Conf

	logger := logger.New(c.Server.Debug)

	if err := envdecode.StrictDecode(&c); err != nil {
		logger.Fatal().Err(err).Msg("Server failed startup")
		// log.Fatalf("Failed to decode: %s", err)
	}

	return &c
}
