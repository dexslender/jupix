package commands

import (
	useractions "github.com/dexslender/plane/commands/user_actions"
	"github.com/dexslender/plane/util"
)

var Commands = []util.JCommand{
	new(Ping),
	new(useractions.GetUserId),
}
