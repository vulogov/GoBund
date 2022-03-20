package bund

import (
	"github.com/pieterclaerhout/go-log"

	// "github.com/vulogov/Bund/internal/conf"
	"github.com/vulogov/Bund/internal/signal"
)

func Eval() {
	Init()
	log.Debug("[ BUND ] bund.Eval() is reached")
	signal.ExitRequest()
}
