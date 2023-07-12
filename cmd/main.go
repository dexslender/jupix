package main

import (
	"github.com/dexslender/plane/commands"
	"github.com/dexslender/plane/plane"
	"github.com/dexslender/plane/util"
	"github.com/disgoorg/log"
	"github.com/disgoorg/snowflake/v2"
)

func main() {
	l := log.New(log.Ltime | log.Lshortfile)
	c := util.SetupConfig(l, ".")
	if c.GetBool("debug-log-level") {
		l.SetLevel(log.LevelDebug)
	}

	p := plane.New(l, c)

	p.SetupBot()

	if c.GetBool("bot~setup-commands") {
		p.Handler.SetupCommands(
			p.Client,
			snowflake.ID(c.GetInt("bot~guild-id")),
			commands.Commands,
		)
	}

	p.StartNLock()
}
