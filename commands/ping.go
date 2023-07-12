package commands

import (
	"github.com/dexslender/plane/util"
	"github.com/disgoorg/disgo/discord"
)

type Ping struct {
	discord.SlashCommandCreate
}

func (c *Ping) Init() {
	c.Name = "ping"
	c.Description = "Latency of the bot"
}

func (c *Ping) Run(ctx util.CommandCtx) error {
	return ctx.CreateMessage(discord.NewMessageCreateBuilder().
		SetEmbeds(discord.NewEmbedBuilder().
			Build(),
		).
		Build(),
	)
}
