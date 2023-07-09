package useractions

import (
	"github.com/dexslender/plane/utils"
	"github.com/disgoorg/disgo/discord"
)

type GetUserId struct {
	discord.UserCommandCreate
}

func (c *GetUserId) Init() {
	c.Name = "Get ID"
}

func (c *GetUserId) Run(ctx utils.CommandCtx) error {
	return ctx.CreateMessage(discord.MessageCreate{Content: ctx.User().ID.String()})
}
