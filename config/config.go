package config

import (
	"time"

	"cloud-native/util/logger"

	"github.com/joeshaw/envdecode"
)

type Conf struct {
	Server serverConf
}

type serverConf struct {
	Port         int           `env:"SERVER_PORT,required"`
	Timeoutread  time.Duration `env:"SERVER_TIMEOUT_READ,required"`
	Timeoutwrite time.Duration `env:"SERVER_TIMEOUT_WRITE,required"`
	Timeoutidle  time.Duration `env:"SERVER_TIMEOUT_IDLE,required"`
	Debug        bool          `env:"DEBUG,required"`
}

func AppConf() *Conf {
	var c Conf

	logger := logger.New(true)

	if err := envdecode.StrictDecode(&c); err != nil {
		logger.Fatal().Err(err).Msg("Server failed startup")
		// log.Fatalf("Failed to decode: %s", err)
	}

	return &c
}
