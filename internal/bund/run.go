package bund

import (
	"github.com/pieterclaerhout/go-log"
	tc "github.com/vulogov/ThreadComputation"
	"github.com/vulogov/Bund/internal/conf"
	"github.com/vulogov/Bund/internal/signal"
)

func RunFile(core *tc.TCstate, name string) {
	log.Debugf("Running: %v", name)
	code, err := tc.ReadFile(name)
	if err != nil {
		log.Fatalf("Error loading file: %v", err)
	}
	core.Eval(code)
}

func Run() {
	Init()
	log.Debug("[ BUND ] bund.Run() is reached")
	if *conf.CDebug {
		log.Info("BUND core debug is on")
		tc.SetVariable("tc.Debuglevel", "debug")
		log.Infof("[ BUND ] core version: %v", tc.VERSION)
	} else {
		log.Debug("BUND core debug is off")
		tc.SetVariable("tc.Debuglevel", "info")
		log.Debugf("[ BUND ] core version: %v", tc.VERSION)
	}
	core := tc.Init()
	for _, f := range *conf.Scripts {
		RunFile(core, f)
	}
	signal.ExitRequest()
}
