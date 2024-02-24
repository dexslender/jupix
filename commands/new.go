package commands

import (
	"github.com/dexslender/jupix/util"
	"github.com/disgoorg/disgo/discord"
)

type NewCmd struct {
	discord.SlashCommandCreate
}

func (c *NewCmd) Init(r util.JHRegister) {
	c.Name = "new"
	c.Description = "Create a new custom command"

	r.Modal("command-maker", func(ctx *util.ModalCtx) error {
		// o := util.Parse(ctx.Data.Text("code"))
		// for _, order := range o {
		// 	order(*ctx)
		// }
		return nil
		// ctx.CreateMessage(discord.NewMessageCreateBuilder().
		// 	SetContent("# Okay, but not working for now...").
		// 	SetEphemeral(true).
		// 	Build(),
		// )
	})
}

func (c *NewCmd) Run(ctx *util.JContext) error {
	return ctx.Modal(discord.NewModalCreateBuilder().
		SetCustomID("command-maker").
		SetTitle("Command Maker").
		// AddActionRow(discord.TextInputComponent{
		// 	CustomID:    "name",
		// 	Style:       discord.TextInputStyleShort,
		// 	Label:       "Name",
		// 	Placeholder: "ping, hello, greet",
		// }).
		AddActionRow(discord.TextInputComponent{
			CustomID: "code",
			Style:    discord.TextInputStyleParagraph,
			Label:    "Code",
			// Placeholder: "> hello",
			Value: "> ",
		}).
		Build(),
	)
}
func (c *NewCmd) Error(ctx *util.JContext, err error) {}
