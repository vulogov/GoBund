package bund

import (
	"github.com/pieterclaerhout/go-log"

	// "github.com/vulogov/Bund/internal/conf"
	"github.com/vulogov/Bund/internal/signal"
)

func Run() {
	Init()
	log.Debug("[ BUND ] bund.Run() is reached")

	signal.ExitRequest()
}
