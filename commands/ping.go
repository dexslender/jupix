package commands

import (
	"github.com/dexslender/plane/utils"
	"github.com/disgoorg/disgo/discord"
)

var _ utils.Command = (*Ping)(nil)

type Ping struct {
	discord.SlashCommandCreate
}

func (c *Ping) Init() {
	c.Name = "ping"
	c.Description = "Latency of the bot"
}

func (c *Ping) Run(ctx utils.CommandCtx) error {
	return ctx.CreateMessage(discord.MessageCreate{Content: "Hello!"})
}
