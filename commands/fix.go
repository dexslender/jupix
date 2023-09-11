package commands

import (
	"github.com/dexslender/jupix/util"
	"github.com/disgoorg/disgo/discord"
	"github.com/disgoorg/json"
)

type FixCMD struct {
	discord.SlashCommandCreate
}

func (c *FixCMD) Init() {
	c.Name = "fix"
	c.Description = "Fix things with this command"
	c.DMPermission = json.Ptr(false)
	c.DefaultMemberPermissions = json.NewNullablePtr(discord.PermissionAdministrator)
	c.Options = []discord.ApplicationCommandOption{
		discord.ApplicationCommandOptionString{
			Required:    true,
			Name:        "target",
			Description: "Supported options",
			Choices: []discord.ApplicationCommandOptionChoiceString{
				{
					Name:  "Rol: Members",
					Value: "rol-members",
				},
			},
		},
	}
}

func (c *FixCMD) Run(ctx *util.JContext) error {
	opt := ctx.SlashCommandInteractionData().String("target")
	switch opt {
	case "rol-members":
		if err := ctx.DeferCreateMessage(false); err != nil {
			return err
		}
		go func(ctx *util.JContext) {
			var mc int
			if g, err := ctx.Client().Rest().GetGuild(*ctx.GuildID(), true); err != nil {
				mc = g.MemberCount
			}
			ctx.Client().Rest().GetMembers(
				*ctx.GuildID(),
				mc,
				0,
			)
		}(ctx)
		return nil
	default:
		return ctx.CreateMessage(discord.NewMessageCreateBuilder().
			SetEphemeral(true).
			SetContent("Unknown option").
			Build(),
		)
	}
}
