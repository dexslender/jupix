package util

import (
	"github.com/disgoorg/disgo/bot"
	"github.com/disgoorg/disgo/discord"
	"github.com/disgoorg/disgo/events"
	"github.com/disgoorg/log"
)

func NewIHandler() *JIHandler {
	return &JIHandler{
		commands: make(map[string]JCommand),
	}
}

type JIHandler struct {
	Log      log.Logger
	Config   Config
	commands map[string]JCommand
}

func (h *JIHandler) OnEvent(event bot.Event) {
	e, ok := event.(*events.InteractionCreate)
	if !ok {
		return
	}

	switch i := e.Interaction.(type) {
	case discord.ApplicationCommandInteraction:
		h.Log.Info("Received interaction: ", i.Data.CommandName())
		if c, ok := h.commands[i.Data.CommandName()]; ok {
			ctx := &JContext{
				events.ApplicationCommandInteractionCreate{
					GenericEvent:                  e.GenericEvent,
					ApplicationCommandInteraction: e.Interaction.(discord.ApplicationCommandInteraction),
					Respond:                       e.Respond,
				},
				h.Log,
			}
			if err := c.Run(ctx); err != nil {
				h.Log.Error("Handler error catched: ", err)
			}
		}

		// case discord.AutocompleteInteraction:
		// case discord.ComponentInteraction:
		// case discord.ModalSubmitInteraction:
	default:
		h.Log.Warnf("Unhandled interaction")
	}
}

func (h *JIHandler) WithLogger(l log.Logger) *JIHandler {
	h.Log = l
	return h
}

func (h *JIHandler) WithConfig(c Config) *JIHandler {
	h.Config = c
	return h
}

func (h *JIHandler) SetupCommands(client bot.Client, commands []JCommand) {
	cmds := []discord.ApplicationCommandCreate{}

	for _, c := range commands {
		c.Init()
		cmds = append(cmds, c)
		h.commands[c.CommandName()] = c
	}

	var (
		reg []discord.ApplicationCommand
		err error
	)

	if h.Config.Bot.SetupCommands {
		if h.Config.Bot.GuildId != 0 {
			reg, err = client.Rest().SetGuildCommands(
				client.ApplicationID(),
				h.Config.Bot.GuildId,
				cmds,
			)
		} else if h.Config.Bot.GlobalCommands {
			reg, err = client.Rest().SetGlobalCommands(
				client.ApplicationID(),
				cmds,
			)
		}
	}

	if err != nil {
		h.Log.Error("Error in commands setup: ", err)
	} else {
		h.Log.Infof("Registered %d commands", len(reg))
	}
}

// func (h *JIHandler) Component() {

// }
