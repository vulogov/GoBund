package bund

import (
	"github.com/pieterclaerhout/go-log"

	"github.com/vulogov/Bund/internal/conf"
	"github.com/vulogov/Bund/internal/signal"
)

func Run() {
	Init()
	log.Debug("[ BUND ] bund.Run() is reached")
	for _, f := range *conf.Scripts {
		log.Debugf("%v", f)
	}
	signal.ExitRequest()
}
