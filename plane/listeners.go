package plane

import (
	"github.com/disgoorg/disgo/events"
)

var listeners = &events.ListenerAdapter{
	OnReady: func(event *events.Ready) {
		_p.Log.Infof("Logged in as '@%s'", event.User.Username)
	},
}
