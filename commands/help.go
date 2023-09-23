package commands

import (
	"github.com/dexslender/jupix/util"
	"github.com/disgoorg/disgo/discord"
)

type Help struct {
	discord.SlashCommandCreate
}

func (c *Help) Init() {
	c.Name = "help"
	c.Description = "😉・Do you need help?"
	c.DescriptionLocalizations = map[discord.Locale]string{
		discord.LocaleSpanishES: "😉・Necesitas ayuda?",
	}
}
func (c *Help) Run(ctx *util.JContext) error {
	return ctx.CreateMessage(discord.NewMessageCreateBuilder().
		SetContent("# ✨・Comming soon!").
		SetEphemeral(true).
		Build(),
	)
}
func (c *Help) Error(ctx *util.JContext, err error) {}
