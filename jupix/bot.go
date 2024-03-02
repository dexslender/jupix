package jupix

import (
	"context"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/charmbracelet/log"
	"github.com/dexslender/jupix/commands"
	"github.com/dexslender/jupix/util"
	"github.com/disgoorg/disgo"
	"github.com/disgoorg/disgo/bot"
	"github.com/disgoorg/disgo/discord"
	"github.com/disgoorg/disgo/gateway"
	"github.com/disgoorg/disgo/handler"
	"github.com/disgoorg/snowflake/v2"
)

func New(l *log.Logger, c util.Config) *Jupix {
	return &Jupix{
		Config: c,
		Log:    l,
	}
}

type Jupix struct {
	bot.Client
	Config   util.Config
	Log      *log.Logger
	Handler  handler.Router
	PUpdater util.PUpdater
}

func (j *Jupix) SetupBot() {
	var err error

	j.PUpdater = util.PUpdater{
		Config: j.Config,
		Log:    j.Log,
	}

	j.Handler = handler.New()

	if j.Client, err = disgo.New(
		j.Config.Bot.Token,
		bot.WithGatewayConfigOpts(
			func(cfg *gateway.Config) {
				j.PUpdater.Setup(cfg)
				if j.Config.Bot.MobileOs {
					cfg.Browser = "Discord Android"
				}
			},
			gateway.WithCompress(true),
			gateway.WithIntents(
				gateway.IntentsNonPrivileged|
					gateway.IntentGuildMembers,
			),
		),
		bot.WithLogger(slog.New(j.Log)),
		bot.WithEventListeners(listeners(j), j.Handler),
	); err != nil {
		j.Log.Fatal("Client error: ", err)
	}

	cmd := new(commands.Ping)
	cmd.Init(nil)
	err = handler.SyncCommands(j, []discord.ApplicationCommandCreate{cmd}, []snowflake.ID{j.Config.Bot.GuildId})
	if err != nil {
		j.Log.Error("failed to sync commands", "err", err)
	}

	j.Handler.Command("/ping", func(e *handler.CommandEvent) error {
		return e.CreateMessage(discord.MessageCreate{Content: "Hello"})
	})
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
