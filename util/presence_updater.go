package util

import (
	"context"
	"time"

	"github.com/disgoorg/disgo/bot"
	"github.com/disgoorg/disgo/discord"
	"github.com/disgoorg/disgo/gateway"
	"github.com/disgoorg/log"
)

type PUpdater struct {
	Config  Config
	Log     log.Logger
	current int
}

func ResolvePresence(p Presence) *gateway.MessageDataPresenceUpdate {
	var (
		status   discord.OnlineStatus
		activity discord.Activity
	)

	if p.Name == "" && p.State == "" {
		return nil
	}

	if p.Name != "" {
		activity.Name = p.Name
	}
	if p.State != "" {
		activity.State = &p.State
		if activity.Name == "" {
			activity.Name = p.State
		}
	}

	if p.Status == "" {
		status = "online"
	} else {
		status = p.Status
	}

	switch p.Type {
	case "watching", "watch":
		activity.Type = discord.ActivityTypeWatching
	case "listening", "listen":
		activity.Type = discord.ActivityTypeListening
	case "game", "playing":
		activity.Type = discord.ActivityTypeGame
	case "competing":
		activity.Type = discord.ActivityTypeCompeting
	case "streaming", "stream":
		activity.Type = discord.ActivityTypeStreaming
	case "custom":
		activity.Type = discord.ActivityTypeCustom
	default:
		activity.Type = discord.ActivityTypeGame
	}

	if activity.Type == discord.ActivityTypeStreaming && p.URL != "" {
		activity.URL = &p.URL
	}

	return &gateway.MessageDataPresenceUpdate{
		Status:     status,
		Activities: []discord.Activity{activity},
	}
}

func (pu *PUpdater) Setup(gc *gateway.Config) {
	if pu.Config.PresenceUpdater.Enabled &&
		len(pu.Config.PresenceUpdater.Presences) >= 1 {
		gc.Presence = pu.Next()
	} else {
		pu.Log.Info("presence updater disabled")
	}
}

func (pu *PUpdater) StartUpdater(c bot.Client) {
	if !pu.Config.PresenceUpdater.Enabled ||
		len(pu.Config.PresenceUpdater.Presences) <= 1 {
		return
	}
	ticker := time.NewTicker(pu.Config.PresenceUpdater.Delay)
	for range ticker.C {
		p := pu.Next()
		if p == nil {
			continue
		}
		pu.Log.Debugf("updating bot presence: index:%d", pu.current)
		c.Gateway().Presence().Activities = p.Activities
		c.Gateway().Presence().Status = p.Status
		if err := c.SetPresence(context.Background()); err != nil {
			pu.Log.Error("failed to send presence data to gateway: ", err)
		}
	}
}

func (pu *PUpdater) Next() (data *gateway.MessageDataPresenceUpdate) {
	size := len(pu.Config.PresenceUpdater.Presences)

	p := ResolvePresence(pu.Config.PresenceUpdater.Presences[pu.current])
	if p != nil {
		data = p
	}

	if pu.current >= size-1 {
		pu.current = 0
	} else {
		pu.current++
	}
	return
}
