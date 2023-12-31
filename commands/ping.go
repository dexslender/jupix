package commands

import (
	"time"

	"github.com/dexslender/jupix/util"
	"github.com/disgoorg/disgo/discord"
)

type Ping struct {
	discord.SlashCommandCreate
}

func (c *Ping) Init(util.JHRegister) {
	c.Name = "ping"
	c.Description = "Latency of the bot"
}

func (c *Ping) Run(ctx *util.JContext) error {
	s := time.Now()
	if err := ctx.DeferCreateMessage(false); err != nil {
		return err
	}
	d := time.Since(s)
	_, err := ctx.UpdateInteractionResponse(discord.NewMessageUpdateBuilder().
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
	return err
}

func (t *Ping) Error(ctx *util.JContext, err error) {

}
