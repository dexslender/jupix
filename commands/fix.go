package commands

import (
	"fmt"
	"log/slog"

	"github.com/dexslender/jupix/util"
	"github.com/disgoorg/disgo/discord"
	"github.com/disgoorg/json"
)

type FixCMD struct {
	discord.SlashCommandCreate
}

func (c *FixCMD) Init(util.JHRegister) {
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

// TODO - Organize
func (c *FixCMD) Run(ctx *util.JContext) error {
	opt := ctx.SlashCommandInteractionData().String("target")
	switch opt {
	case "rol-members":
		return OnRolMembers(ctx)
	default:
		return ctx.CreateMessage(discord.NewMessageCreateBuilder().
			SetEphemeral(true).
			SetContent("Unknown option").
			Build(),
		)
	}
}

func OnRolMembers(ctx *util.JContext) error {
	if err := ctx.DeferCreateMessage(false); err != nil {
		return err
	}

	log_err := func(err error) {
		if err != nil {
			ctx.Log.Error(err)
		}
	}

	go func(ctx *util.JContext) {
		var mc int
		if g, err := ctx.
			Client().
			Rest().
			GetGuildPreview(*ctx.GuildID()); err != nil {
			_, err := ctx.UpdateInteractionResponse(discord.NewMessageUpdateBuilder().
				SetContentf("Err fetch guild: %s", err).
				Build(),
			)
			log_err(err)
		} else {
			mc = *g.ApproximateMemberCount
			ctx.Client().Logger().Info("members", slog.Int("count", *g.ApproximateMemberCount))
		}

		if ms, err := ctx.
			Client().
			Rest().
			GetMembers(
				*ctx.GuildID(),
				mc,
				0,
			); err != nil {
			_, err := ctx.UpdateInteractionResponse(discord.NewMessageUpdateBuilder().
				SetContentf("Err fetch guild members: %s", err).
				Build(),
			)
			log_err(err)
		} else {
			membs := "```js"
			for i, m := range ms {
				membs += fmt.Sprintf("\n%d- %d", i+1, m.User.ID)
				if len(ms)-1 == i {
					membs += "\n```"
				}
			}
			_, err := ctx.UpdateInteractionResponse(discord.MessageUpdate{
				Content: json.Ptr(membs),
			})
			log_err(err)
		}
	}(ctx)
	return nil
}

func (t *FixCMD) Error(ctx *util.JContext, err error) {

}
