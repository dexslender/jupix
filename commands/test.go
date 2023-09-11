package commands

import (
	"github.com/dexslender/jupix/util"
	"github.com/disgoorg/disgo/discord"
)

type TestCommand struct {
	discord.SlashCommandCreate
}

func (t *TestCommand) Init() {
	panic("not implemented") // TODO: Implement
}

// Before(ctx *JContext) bool
// RunError(ctx *JContext, err error)
// After(ctx *JContext)
func (t *TestCommand) Run(ctx *util.JContext) error {
	panic("not implemented") // TODO: Implement
}

