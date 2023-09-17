package commands

import (
	"github.com/dexslender/jupix/util"
	"github.com/disgoorg/disgo/discord"
)

type CommandTest struct {
	discord.ApplicationCommandCreate
}

func (c *CommandTest) Init() {}
func (c *CommandTest) Run(ctx *util.JContext) error {
	return nil
}
func (c *CommandTest) Error(ctx *util.JContext, err error) {}
