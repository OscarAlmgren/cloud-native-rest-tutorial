package config

import (
	"log"
	"time"

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
}

func AppConf() *Conf {
	var c Conf

	if err := envdecode.StrictDecode(&c); err != nil {
		log.Fatalf("Failed to decode: %s", err)
	}

	return &c
}
