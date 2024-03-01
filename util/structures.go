package util

import (
	"github.com/disgoorg/disgo/events"
	"github.com/charmbracelet/log"
)

type ComponentHandle func(ctx *ComponentCtx) error

type ComponentCtx struct {
	events.ComponentInteractionCreate
	Log *log.Logger
}

type ModalHandle func(ctx *ModalCtx) error

type ModalCtx struct {
	events.ModalSubmitInteractionCreate
	Log *log.Logger
}

// return the result?
type AutocompleteHandle func(ctx *AutocompleteCtx) error

type AutocompleteCtx struct {
	events.AutocompleteInteractionCreate
	Log *log.Logger
}

type JComponent struct{}

type JModal struct{}

type JAutocomplete struct{}
