package plane

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/disgoorg/disgo"
	"github.com/disgoorg/disgo/bot"
	"github.com/disgoorg/disgo/gateway"
	"github.com/disgoorg/log"
	"github.com/spf13/viper"
)

var _p Plane

func New(l log.Logger, c *viper.Viper) *Plane {
	_p = Plane{
		Config: c,
		Log:    l,
	}

	return &_p
}

type Plane struct {
	bot.Client
	Config *viper.Viper
	Log    log.Logger
}

func (p *Plane) SetupBot() {
	var err error
	if p.Client, err = disgo.New(
		p.Config.GetString("bot~token"),
		bot.WithGatewayConfigOpts(
			func(cfg *gateway.Config) {
				if p.Config.GetBool("bot~mobile-os") {
					cfg.Browser = "Discord Android"
				}
			},
			gateway.WithCompress(true),
			gateway.WithIntents(gateway.IntentsNonPrivileged),
		),
		bot.WithLogger(p.Log),
		bot.WithEventListeners(listeners),
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
