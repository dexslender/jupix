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
	Conf    Config
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

func (pu *PUpdater) InitUpdater(c bot.Client) {
	ticker := time.NewTicker(pu.Conf.Bot.PresenceInterval)
	for range ticker.C {
		pu.Log.Debug("Updating presence...")
		presence := pu.Next()
		if presence == nil {
			continue
		}
		c.Gateway().Presence().Activities = presence.Activities
		c.Gateway().Presence().Status = presence.Status
		if err := c.SetPresence(context.Background()); err != nil {
			pu.Log.Error("Error in presence handler: ", err)
		}
	}
}

func (pu *PUpdater) Next() (data *gateway.MessageDataPresenceUpdate) {
	size := len(pu.Conf.Presences)

	if p := ResolvePresence(pu.Conf.Presences[pu.current]); p != nil {
		data = p
	}

	if pu.current >= size-1 {
		pu.current = 0
	} else {
		pu.current++
	}
	return
}
