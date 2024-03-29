package bund

import (
	"fmt"
	"strings"
	"github.com/lrita/cmap"
	"github.com/peterh/liner"
	"github.com/pieterclaerhout/go-log"
	tc "github.com/vulogov/ThreadComputation"
	"github.com/vulogov/Bund/internal/conf"
	"github.com/vulogov/Bund/internal/banner"
)

var (
	shellCmd cmap.Cmap
	commands = []string{
		".version", ".exit", ".stack", ".last",
	}
	PROMPT = "[ BUND ] "
)

func Shell() {
	Init()
	banner.PrintBanner(fmt.Sprintf("[ BUND %v ]", conf.EVersion))
	log.Info("For exit, type: .exit")
	log.Debug("[ BUND ] bund.Shell() is reached")
	line := liner.NewLiner()
	defer line.Close()
	line.SetCtrlCAborts(true)

	line.SetCompleter(func(line string) (c []string) {
		for _, n := range commands {
			if strings.HasPrefix(n, strings.ToLower(line)) {
				c = append(c, n)
			}
		}
		return
	})
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

	out:
	for {
		if cmd, err := line.Prompt(PROMPT); err == nil {
			cmd = strings.Trim(cmd, "\n \t\r")
			line.AppendHistory(cmd)
			log.Debugf("shell get: %v", cmd)
			switch cmd {
			case ".exit":
				log.Debug("Exiting")
				break out
			default:
				if IsShellCommand(cmd) {
					log.Debugf("Running shell command: %v", cmd)
					RunShellCommand(cmd, core)
				} else {
					log.Debug("Executing in ThreadComputation")
					core.Eval(cmd)
					ShellDisplayResult(core, false)
				}
			}
		} else if err == liner.ErrPromptAborted {
			log.Debug("Aborted")
			break
		} else {
			log.Debugf("Error reading line: %v", err)
		}
	}
}
