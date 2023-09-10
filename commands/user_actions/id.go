package useractions

import (
	"github.com/dexslender/plane/util"
	"github.com/disgoorg/disgo/discord"
)

type GetUserId struct {
	discord.UserCommandCreate
}

func (c *GetUserId) Init() {
	c.Name = "Get ID"
}

func (c *GetUserId) Run(ctx util.CommandCtx) error {
	return ctx.CreateMessage(discord.NewMessageCreateBuilder().
		SetEphemeral(true).
		SetContentf(
			"%s%s%s",
			"_`",
			ctx.User().ID.String(),
			"`_",
		).
		Build(),
	)
}