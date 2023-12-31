package commands

import (
	useractions "github.com/dexslender/jupix/commands/user_actions"
	"github.com/dexslender/jupix/util"
)

var Commands = []util.JCommand{
	new(Ping),
	new(useractions.GetUserId),
	new(FixCMD),
	new(Help),
	new(CommandTest),
	new(NewCmd),
}
