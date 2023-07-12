package commands

import (
	"time"

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
	s := time.Now()
	if err := ctx.DeferCreateMessage(false); err != nil {
		return err
	}
	d := time.Since(s)

	return ctx.CreateMessage(discord.NewMessageCreateBuilder().
		SetEmbeds(discord.NewEmbedBuilder().
			SetColor(util.DDark).
			SetAuthorName("Pong!").
			AddField(
				"Gateway",
				ctx.Client().Gateway().Latency().String(),
				false,
			).
			AddField(
				"Rest",
				d.String(),
				false,
			).
			Build(),
		).
		Build(),
	)
}
