package main

import (
	"os"

	"github.com/charmbracelet/log"
	"github.com/dexslender/jupix/jupix"
	"github.com/dexslender/jupix/util"
)

func main() {
	l := log.NewWithOptions(os.Stderr, log.Options{
		ReportTimestamp: true,
	})
	c := util.LoadConfig(l, "botconfig.yaml", "JUPIX")
	if c.Bot.DebugLog {
		l.SetLevel(log.DebugLevel)
	}

	p := jupix.New(l, c)
	p.SetupBot()
	p.StartNLock()
}
