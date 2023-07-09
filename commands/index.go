package commands

import (
	useractions "github.com/dexslender/plane/commands/user_actions"
	"github.com/dexslender/plane/utils"
)

var Commands = []utils.Command{
	new(Ping),
	new(useractions.GetUserId),
}
