package plane

import (
	"github.com/disgoorg/disgo/discord"
	"github.com/disgoorg/disgo/events"
	"github.com/disgoorg/json"
	"github.com/disgoorg/snowflake/v2"
)

var listeners = &events.ListenerAdapter{
	OnReady: func(event *events.Ready) {
		_p.Log.Infof("Logged in as '@%s'", event.User.Username)

		if _p.Config.GetBool("bot~setup-commands") {
			if _, err := _p.Rest().SetGuildCommands(
				_p.ApplicationID(),
				snowflake.ID(_p.Config.GetInt("bot~guild-id")),
				[]discord.ApplicationCommandCreate{
					discord.SlashCommandCreate{
						Name:                     "fix",
						Description:              "Fix some thing",
						DefaultMemberPermissions: json.NewNullablePtr[discord.Permissions](discord.PermissionAdministrator),
						Options: []discord.ApplicationCommandOption{
							discord.ApplicationCommandOptionString{
								Name:        "target",
								Description: "Supported options",
								Required:    true,
								Choices: []discord.ApplicationCommandOptionChoiceString{
									{
										Name:  "Member Role",
										Value: "member-role",
									},
								},
							},
						},
					},
				},
			); err != nil {
				_p.Log.Error("Error while registering slash commands: ", err)
			}
		}
	},
}
