package main

import (
	"github.com/rs/zerolog/log"
	"time"
)

func main() {
	log.Info().Msg("reflector-server is starting...")
	for {
		time.Sleep(1 * time.Second)
		log.Info().Msg("holding...")
	}
}
