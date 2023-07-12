package util

import (
	"github.com/disgoorg/disgo/bot"
	"github.com/disgoorg/disgo/discord"
	"github.com/disgoorg/disgo/events"
	"github.com/disgoorg/log"
	"github.com/disgoorg/snowflake/v2"
)

type Command interface {
	discord.ApplicationCommandCreate
	Init()
	Run(ctx CommandCtx) error
}

type CommandCtx struct {
	events.ApplicationCommandInteractionCreate
}

func NewHandler() *Handler {
	return &Handler{
		commands: make(map[string]Command),
	}
}

type Handler struct {
	Log      log.Logger
	commands map[string]Command
}

/*
For example propourses
ApplicationCommandInteractionCreate: &events.ApplicationCommandInteractionCreate{
	GenericEvent:                  event.GenericEvent,
	ApplicationCommandInteraction: event.Interaction.(discord.ApplicationCommandInteraction),
	Respond:                       event.Respond,
}
*/

func (h *Handler) OnEvent(event bot.Event) {
	e, ok := event.(*events.InteractionCreate)
	if !ok {
		return
	}

	switch i := e.Interaction.(type) {
	case discord.ApplicationCommandInteraction:
		h.Log.Info("Received interaction: ", i.Data.CommandName())
		if c, ok := h.commands[i.Data.CommandName()]; ok {
			ctx := CommandCtx{
				events.ApplicationCommandInteractionCreate{
					GenericEvent:                  e.GenericEvent,
					ApplicationCommandInteraction: e.Interaction.(discord.ApplicationCommandInteraction),
					Respond:                       e.Respond,
				},
			}

			c.Run(ctx)
		}

		// case discord.AutocompleteInteraction:
		// case discord.ComponentInteraction:
		// case discord.ModalSubmitInteraction:
	default:
		h.Log.Info("Unhandled interaction")
	}
}

func (h *Handler) WithLogger(l log.Logger) *Handler {
	h.Log = l
	return h
}

func (h *Handler) SetupCommands(c bot.Client, guildId snowflake.ID, commands []Command) {
	if guildId == 0 {
		return
	}

	cmds := []discord.ApplicationCommandCreate{}

	for _, c := range commands {
		c.Init()
		cmds = append(cmds, c)
		h.commands[c.CommandName()] = c
	}

	if reg, err := c.Rest().SetGuildCommands(
		c.ApplicationID(),
		guildId,
		cmds,
	); err != nil {
		h.Log.Error("Error in command setup: ", err)
	} else {
		h.Log.Infof("Registered %d commands", len(reg))
	}
}
