package useractions

import (
	"github.com/dexslender/jupix/util"
	"github.com/disgoorg/disgo/discord"
)

type GetUserId struct {
	discord.UserCommandCreate
}

func (c *GetUserId) Init() {
	c.Name = "Get ID"
}

func (c *GetUserId) Run(ctx *util.JContext) error {
	return ctx.CreateMessage(discord.NewMessageCreateBuilder().
		SetEphemeral(true).
		SetContentf(
			"%s\n%s%s",
			"```go",
			ctx.User().ID.String(),
			"```",
		).
		Build(),
	)
}
