package util

import (
	"github.com/disgoorg/disgo/bot"
	"github.com/disgoorg/disgo/discord"
	"github.com/disgoorg/disgo/events"
	"github.com/disgoorg/disgo/rest"
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

func (e *CommandCtx) GetInteractionResponse(opts ...rest.RequestOpt) (*discord.Message, error) {
	return e.Client().Rest().GetInteractionResponse(e.ApplicationID(), e.Token(), opts...)
}

func (e *CommandCtx) UpdateInteractionResponse(messageUpdate discord.MessageUpdate, opts ...rest.RequestOpt) (*discord.Message, error) {
	return e.Client().Rest().UpdateInteractionResponse(e.ApplicationID(), e.Token(), messageUpdate, opts...)
}

func (e *CommandCtx) DeleteInteractionResponse(opts ...rest.RequestOpt) error {
	return e.Client().Rest().DeleteInteractionResponse(e.ApplicationID(), e.Token(), opts...)
}

func (e *CommandCtx) GetFollowupMessage(messageID snowflake.ID, opts ...rest.RequestOpt) (*discord.Message, error) {
	return e.Client().Rest().GetFollowupMessage(e.ApplicationID(), e.Token(), messageID, opts...)
}

func (e *CommandCtx) CreateFollowupMessage(messageCreate discord.MessageCreate, opts ...rest.RequestOpt) (*discord.Message, error) {
	return e.Client().Rest().CreateFollowupMessage(e.ApplicationID(), e.Token(), messageCreate, opts...)
}

func (e *CommandCtx) UpdateFollowupMessage(messageID snowflake.ID, messageUpdate discord.MessageUpdate, opts ...rest.RequestOpt) (*discord.Message, error) {
	return e.Client().Rest().UpdateFollowupMessage(e.ApplicationID(), e.Token(), messageID, messageUpdate, opts...)
}

func (e *CommandCtx) DeleteFollowupMessage(messageID snowflake.ID, opts ...rest.RequestOpt) error {
	return e.Client().Rest().DeleteFollowupMessage(e.ApplicationID(), e.Token(), messageID, opts...)
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

			if err := c.Run(ctx); err != nil {
				h.Log.Error("Handler error catched: ", err)
			}
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
		h.Log.Error("Error in commands setup: ", err)
	} else {
		h.Log.Infof("Registered %d commands", len(reg))
	}
}

func (h *Handler) Component() {

}
