package commands

import (
	"github.com/dexslender/jupix/util"
	"github.com/disgoorg/disgo/discord"
)

type CommandTest struct {
	discord.SlashCommandCreate
}

func (c *CommandTest) Init(util.JHRegister) {
	c.Name = "test"
	c.Description = "A test command :D"
}
func (c *CommandTest) Run(ctx *util.JContext) error {
	// select_user := discord.NewUserSelectMenu("select_target", "Select three users here!").
	// 	WithMaxValues(3)

	return ctx.CreateMessage(discord.NewMessageCreateBuilder().
		SetEphemeral(true).
		SetContent("no specified.").
		Build(),
	)
}
func (c *CommandTest) Error(ctx *util.JContext, err error) {}
