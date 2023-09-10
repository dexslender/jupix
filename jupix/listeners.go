package jupix

import (
	"github.com/disgoorg/disgo/events"
)

var listeners = &events.ListenerAdapter{
	OnReady: func(event *events.Ready) {
		go jx.Presences.InitUpdater(event.Client())
		jx.Log.Infof("Logged in as '@%s'", event.User.Username)
	},
}
