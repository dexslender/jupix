package jupix

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/dexslender/jupix/commands"
	"github.com/dexslender/jupix/util"
	"github.com/disgoorg/disgo"
	"github.com/disgoorg/disgo/bot"
	"github.com/disgoorg/disgo/gateway"
	"github.com/disgoorg/log"
)

var jx Jupix

func New(l log.Logger, c util.Config) *Jupix {
	jx = Jupix{
		Config: c,
		Log:    l,
	}
	return &jx
}

type Jupix struct {
	bot.Client
	Config    util.Config
	Log       log.Logger
	Handler   *util.JIHandler
	Presences util.PUpdater
}

func (j *Jupix) SetupBot() {
	var err error

	j.Presences = util.PUpdater{
		Conf: j.Config,
		Log:  j.Log,
	}

	j.Handler = util.NewHandler().
		WithLogger(j.Log).
		WithConfig(j.Config)

	if j.Client, err = disgo.New(
		j.Config.Bot.Token,
		bot.WithGatewayConfigOpts(
			func(cfg *gateway.Config) {
				if len(j.Config.Presences) >= 1 {
					cfg.Presence = j.Presences.Next()
				}
				if j.Config.Bot.MobileOs {
					cfg.Browser = "Discord Android"
				}
			},
			gateway.WithCompress(true),
			gateway.WithIntents(gateway.IntentsNonPrivileged),
		),
		bot.WithLogger(j.Log),
		bot.WithEventListeners(listeners, j.Handler),
	); err != nil {
		j.Log.Fatal("Client error: ", err)
	}

	j.Handler.SetupCommands(j.Client, commands.Commands)
}

func (j *Jupix) StartNLock() {
	ctx, c := context.WithTimeout(context.Background(), time.Second*10)
	defer c()

	defer func() {
		j.Close(ctx)
		j.Log.Info("Client closed, exiting program\n\t(please wait)")
	}()

	if err := j.OpenGateway(ctx); err != nil {
		j.Log.Fatal("Gateway open error: ", err)
	}

	k := make(chan os.Signal, 1)
	signal.Notify(k, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	<-k
}
