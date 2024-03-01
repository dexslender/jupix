package jupix

import (
	"github.com/disgoorg/disgo/bot"
	"github.com/disgoorg/disgo/events"
)

func listeners(jx *Jupix) bot.EventListener {
	return &events.ListenerAdapter{
		OnReady: func(event *events.Ready) {
			go jx.PUpdater.StartUpdater(event.Client())
			jx.Log.Info("Logged in", "username", event.User.Username)
		},
	}
}
