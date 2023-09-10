package util

import (
	"github.com/disgoorg/disgo/discord"
	"github.com/disgoorg/disgo/gateway"
)

type PUpdater struct {
	Conf    *Config
	current int
}

func ResolvePresence(p Presence) *gateway.MessageDataPresenceUpdate {
	var (
		status   discord.OnlineStatus
		activity discord.Activity
	)

	if p.Name == "" {
		return nil
	} else {
		activity.Name = p.Name
	}

	if p.Status == "" {
		status = "online"
	} else {
		status = p.Status
	}

	switch p.Type {
	case "watching":
		activity.Type = discord.ActivityTypeWatching
	case "listening":
		activity.Type = discord.ActivityTypeListening
	case "game", "playing":
		activity.Type = discord.ActivityTypeGame
	case "competing":
		activity.Type = discord.ActivityTypeCompeting
	case "streaming":
		activity.Type = discord.ActivityTypeStreaming
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
