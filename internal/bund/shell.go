package bund

import (
	"strings"

	"github.com/peterh/liner"
	"github.com/pieterclaerhout/go-log"
)

var (
	commands = []string{

	}
)

func Shell() {
	Init()
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
	for {
		if cmd, err := line.Prompt("BUND> "); err == nil {
			log.Debugf("shell get: %v", cmd)
		} else if err == liner.ErrPromptAborted {
			log.Debug("Aborted")
			break
		} else {
			log.Debugf("Error reading line: %v", err)
		}
	}
}
