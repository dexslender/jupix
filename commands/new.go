package commands

import (
	"github.com/dexslender/jupix/util"
	"github.com/disgoorg/disgo/discord"
)

type NewCmd struct {
	discord.SlashCommandCreate
}

func (c *NewCmd) Init(util.JHRegister) {
	c.Name = "new"
	c.Description = "Create a new custom command"
}
func (c *NewCmd) Run(ctx *util.JContext) error {
	return ctx.Modal(discord.NewModalCreateBuilder().
		SetCustomID("command-maker").
		SetTitle("Command Maker").
		AddActionRow(discord.TextInputComponent{
			CustomID:    "name",
			Style:       discord.TextInputStyleShort,
			Label:       "Name",
			Placeholder: "ping, hello, greet",
		}).
		AddActionRow(discord.TextInputComponent{
			CustomID:    "code",
			Style:       discord.TextInputStyleParagraph,
			Label:       "Code",
			Placeholder: "working on it\n>",
		}).
		Build(),
	)
}
func (c *NewCmd) Error(ctx *util.JContext, err error) {}
