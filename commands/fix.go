package commands

import (
	"github.com/dexslender/plane/util"
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

func (c *FixCMD) Run(ctx util.CommandCtx) error {
	return nil
}
