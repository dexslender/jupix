package main

import (
	"github.com/dexslender/jupix/jupix"
	"github.com/dexslender/jupix/util"
	"github.com/disgoorg/log"
)

func main() {
	l := log.New(log.Ltime | log.Lshortfile)
	c := util.LoadConfig(l, "botconfig.yaml", "JUPIX")
	if c.Bot.DebugLog {
		l.SetLevel(log.LevelDebug)
	}

	p := jupix.New(l, c)
	p.SetupBot()
	p.StartNLock()
}
