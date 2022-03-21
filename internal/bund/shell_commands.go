package bund

import (
	"github.com/pieterclaerhout/go-log"
	tc "github.com/vulogov/ThreadComputation"
	"github.com/vulogov/Bund/internal/banner"
)

type ShellCommand func(*tc.TCstate) interface{}

func AddShellCommand(name string, fun ShellCommand) {
	shellCmd.Delete(name)
	shellCmd.Store(name, fun)
}

func RunShellCommand(name string, core *tc.TCstate) {
	if fun, ok := shellCmd.Load(name); ok {
		res := fun.(ShellCommand)(core)
		if res != nil {
			log.Infof("Returned: %v", res)
		}
	} else {
		log.Errorf("Shell command: %v not found", name)
	}
}

func IsShellCommand(name string) bool {
	if _, ok := shellCmd.Load(name); ok {
		return true
	}
	return false
}

func ShellCommandVersion(core *tc.TCstate) interface{} {
	banner.Table(true)
	return nil
}

func init() {
	AddShellCommand(".version", ShellCommandVersion)
}
