package main

import (
	"github.com/dexslender/plane/commands"
	"github.com/dexslender/plane/jupix"
	"github.com/dexslender/plane/util"
	"github.com/disgoorg/log"
)

func main() {
	l := log.New(log.Ltime | log.Lshortfile)
	c := util.LoadConfig(l, "botconfig.yaml")
	if c.Runtime.DebugLog {
		l.SetLevel(log.LevelDebug)
	}

	p := jupix.New(l, c)

	p.SetupBot()

	if c.Bot.SetupCommands {
		p.Handler.SetupCommands(
			p.Client,
			c.Bot.GuildId,
			commands.Commands,
		)
	}

	p.StartNLock()
}
