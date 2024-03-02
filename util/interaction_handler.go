package util

import (
	"github.com/charmbracelet/log"
	"github.com/disgoorg/disgo/bot"
	"github.com/disgoorg/disgo/discord"
	"github.com/disgoorg/disgo/events"
)

var _ JHRegister = (*JIHandler)(nil)

func NewIHandler(l *log.Logger, c Config) *JIHandler {
	return &JIHandler{
		Log:        l,
		Config:     c,
		commands:   make(map[string]JCommand),
		components: make(map[string]ComponentHandle),
		modals:     make(map[string]ModalHandle),
		autocompls: make(map[string]AutocompleteHandle),
	}
}

type JIHandler struct {
	Log        *log.Logger
	Config     Config
	commands   map[string]JCommand
	components map[string]ComponentHandle
	modals     map[string]ModalHandle
	autocompls map[string]AutocompleteHandle
}

type JHRegister interface {
	Component(customID string, handle ComponentHandle)
	Modal(customID string, handle ModalHandle)
	Autocomplete(p string, handle AutocompleteHandle)
}

func (h *JIHandler) OnEvent(event bot.Event) {
	e, ok := event.(*events.InteractionCreate)
	if !ok {
		return
	}

	switch i := e.Interaction.(type) {
	case discord.ApplicationCommandInteraction:
		h.Log.Info("received interaction", "name", i.Data.CommandName())
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
				c.Error(ctx, err)
			}
		}

	case discord.ComponentInteraction:
		h.Log.Info("received component", ComponentInteractionType[i.Data.Type()], i.Data.CustomID())
		if c, ok := h.components[i.Data.CustomID()]; ok {
			ctx := &ComponentCtx{
				events.ComponentInteractionCreate{
					GenericEvent:         e.GenericEvent,
					ComponentInteraction: e.Interaction.(discord.ComponentInteraction),
					Respond:              e.Respond,
				},
				h.Log,
			}
			if err := c(ctx); err != nil {
				h.Log.Error("component interaction error: ", "err", err)
			}
		}

	case discord.ModalSubmitInteraction:
		h.Log.Info("received modal", "custom_id", i.Data.CustomID)
		if m, ok := h.modals[i.Data.CustomID]; ok {
			ctx := &ModalCtx{
				events.ModalSubmitInteractionCreate{
					GenericEvent:           e.GenericEvent,
					ModalSubmitInteraction: e.Interaction.(discord.ModalSubmitInteraction),
					Respond:                e.Respond,
				},
				h.Log,
			}
			if err := m(ctx); err != nil {
				h.Log.Error("modal submit error", "err", err)
			}
		}

	case discord.AutocompleteInteraction:
		h.Log.Info("received autocomplete request", "command", i.Data.CommandName)
		if a, ok := h.autocompls[i.Data.CommandName]; ok {
			ctx := &AutocompleteCtx{
				events.AutocompleteInteractionCreate{
					GenericEvent:            e.GenericEvent,
					AutocompleteInteraction: e.Interaction.(discord.AutocompleteInteraction),
					Respond:                 e.Respond,
				},
				h.Log,
			}
			if err := a(ctx); err != nil {
				h.Log.Error("failed to send autocomplete", "err", err)
			}
		}

	default:
		h.Log.Warnf("unhandled '%s' interaction", InteractionTypeString[i.Type()])
	}
}

func (h *JIHandler) SetupCommands(client bot.Client, commands []JCommand) {
	cmds := []discord.ApplicationCommandCreate{}

	for _, c := range commands {
		c.Init(JHRegister(h))
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
		h.Log.Error("error in commands setup", "err", err)
	} else if h.Config.Bot.SetupCommands {
		h.Log.Infof("registered %d commands", len(reg))
	}
}

func (h *JIHandler) Component(customID string, handle ComponentHandle) {
	h.components[customID] = handle
}

func (h *JIHandler) Modal(customID string, handle ModalHandle) {
	h.modals[customID] = handle
}

func (h *JIHandler) Autocomplete(p string, handle AutocompleteHandle) {

}

// Debug purposes
var ComponentInteractionType = map[discord.ComponentType]string{
	discord.ComponentTypeActionRow:             `action_row`,
	discord.ComponentTypeButton:                `button`,
	discord.ComponentTypeStringSelectMenu:      `string_select_menu`,
	discord.ComponentTypeTextInput:             `text_input`,
	discord.ComponentTypeUserSelectMenu:        `user_select_menu`,
	discord.ComponentTypeRoleSelectMenu:        `role_select_menu`,
	discord.ComponentTypeMentionableSelectMenu: `mentionable_select_menu`,
	discord.ComponentTypeChannelSelectMenu:     `channel_select_menu`,
}

// Debug purposes
var InteractionTypeString = map[discord.InteractionType]string{
	discord.InteractionTypeComponent:          `component`,
	discord.InteractionTypeAutocomplete:       `autocomplete`,
	discord.InteractionTypeModalSubmit:        `modal`,
	discord.InteractionTypeApplicationCommand: `command`,
	discord.InteractionTypePing:               `ping`,
}
