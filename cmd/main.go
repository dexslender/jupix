package main

import (
	"github.com/dexslender/plane/commands"
	"github.com/dexslender/plane/plane"
	"github.com/dexslender/plane/util"
	"github.com/disgoorg/log"
)

func main() {
	l := log.New(log.Ltime | log.Lshortfile)
	c := util.LoadConfig(l, "botconfig.yml")
	if c.Config.LogDebug {
		l.SetLevel(log.LevelDebug)
	}

	p := plane.New(l, c)

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
