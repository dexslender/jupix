package plane

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/dexslender/plane/util"
	"github.com/disgoorg/disgo"
	"github.com/disgoorg/disgo/bot"
	"github.com/disgoorg/disgo/gateway"
	"github.com/disgoorg/log"
)

var _p Plane

func New(l log.Logger, c util.Config) *Plane {
	_p = Plane{
		Config: c,
		Log:    l,
	}

	return &_p
}

type Plane struct {
	bot.Client
	Config  util.Config
	Log     log.Logger
	Handler *util.Handler
}

func (p *Plane) SetupBot() {
	var err error

	p.Handler = util.NewHandler().
		WithLogger(p.Log)

	if p.Client, err = disgo.New(
		p.Config.Bot.Token,
		bot.WithGatewayConfigOpts(
			func(cfg *gateway.Config) {
				if p.Config.Bot.MobileOs {
					cfg.Browser = "Discord Android"
				}
			},
			gateway.WithCompress(true),
			gateway.WithIntents(gateway.IntentsNonPrivileged),
		),
		bot.WithLogger(p.Log),
		bot.WithEventListeners(listeners, p.Handler),
	); err != nil {
		p.Log.Fatal("Client error: ", err)
	}
}

func (p *Plane) StartNLock() {
	ctx, c := context.WithTimeout(context.Background(), time.Second*10)
	defer c()

	defer func() {
		p.Close(ctx)
		p.Log.Info("Client closed, exiting program\n\t(please wait)")
	}()

	if err := p.OpenGateway(ctx); err != nil {
		p.Log.Fatal("Gateway open error: ", err)
	}

	k := make(chan os.Signal, 1)
	signal.Notify(k, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	<-k
}
