package conf

import (
	"fmt"
	"os"
	"time"

	"gopkg.in/alecthomas/kingpin.v2"
)

type filelist []string

var Argv [][]string

func (i *filelist) Set(value string) error {
	_, err := os.Stat(value)
	if os.IsNotExist(err) {
		return fmt.Errorf("Script file '%s' not found", value)
	} else {
		*i = append(*i, value)
		return nil
	}
}

func (i *filelist) String() string {
	return ""
}

func (i *filelist) IsCumulative() bool {
	return true
}

func FileList(s kingpin.Settings) (target *[]string) {
	target = new([]string)
	s.SetValue((*filelist)(target))
	return
}

var (
	seed    = time.Now().UTC().UnixNano()
	App     = kingpin.New("BUND", fmt.Sprintf("[ BUND ] Language that is Functional and Stack-based: %v", BVersion))
	Debug   = App.Flag("debug", "Enable debug mode.").Default("false").Bool()
	CDebug  = App.Flag("core-debug", "Enable core debug mode.").Default("false").Bool()
	Color   = App.Flag("color", "--color : Enable colors on terminal --no-color : Disable colors .").Default("true").Bool()
	VBanner = App.Flag("banner", "Display [ BUND ] banner .").Default("false").Bool()
	Args    = App.Flag("args", "String of arguments passed to a script").String()


	Version = App.Command("version", "Display information about [ BUND ]")
	VTable  = Version.Flag("table", "Display [ BUND ] inner information .").Default("true").Bool()

	Shell      	= App.Command("shell", "Run [ BUND ] in interactive shell")
	ShowSResult = Shell.Flag("result", "Display result of expressions evaluated in [ BUND ] shell").Default("false").Bool()
	SExpr 			= Shell.Arg("expression", "[ BUND ] expression passed to shell.").String()

	Run        	= App.Command("run", "Run BUND in non-interactive mode")
	Scripts    	= Run.Arg("Scripts", "[ BUND ] code to load").Strings()
	ShowRResult = Run.Flag("result", "Display result of scripts execution as it returned by [ BUND ]").Default("false").Bool()

	Eval 				= App.Command("eval", "Evaluate a [ BUND ] expression")
	EStdin  		= Eval.Flag("stdin", "Read [ BUND ] expression from STDIN .").Default("false").Bool()
	Expr 				= Eval.Arg("expression", "[ BUND ] expression.").String()
	ShowEResult = Eval.Flag("result", "Display result of [ BUND ] expression evaluation").Default("false").Bool()
)
