package util

import (
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/disgoorg/disgo/discord"
	"github.com/disgoorg/log"
	"github.com/disgoorg/snowflake/v2"
	"github.com/kkyr/fig"
)

const DEFAULT_CONTENT = `
bot:
#  token: can be used as env-var
  guild: 0
  setup-commands: false
  global-commands: false
  use-mobile-os: true
  debug-log: false

presence-updater:
  multi-presence: false
  presence-interval: 10s
  presences:
    - state: Online now!
      type: custom
      status: online

    # - name:
    #   type:
    #   status:
    #   url:
    #   state:
`

type Config struct {
	Bot struct {
		Token          string       `validate:"required"`
		GuildId        snowflake.ID `fig:"guild" default:"0"`
		SetupCommands  bool         `fig:"setup-commands"`
		GlobalCommands bool         `fig:"global-commands"`
		MobileOs       bool         `fig:"use-mobile-os"`
		DebugLog       bool         `fig:"debug-log"`
	}
	PresenceUpdater struct {
		MultiPresence    bool          `fig:"multi-presence"`
		PresenceInterval time.Duration `fig:"presence-interval" default:"10s"`
		Presences        []Presence
	} `fig:"presence-updater"`
}

type Presence struct {
	Status discord.OnlineStatus
	Name   string
	Type   string
	URL    string
	State  string
}

func LoadConfig(l log.Logger, filename string, env string) (config Config) {
	err := fig.Load(
		&config,
		fig.UseEnv(env),
		fig.File(filename),
	)

	if err != nil {
		if errors.Is(err, fig.ErrFileNotFound) {
			l.Infof("Config file not found, creating in '%s'", fig.DefaultDir)
			if err := WriteConfig(filename); err != nil {
				log.Error("Error while written config file: ", err)
			} else {
				if err := fig.Load(
					&config,
					fig.UseEnv(env),
					fig.File(filename),
				); err != nil {
					l.Error("Config error: ", err)
				}
			}
		} else {
			l.Fatal("Config error: ", err)
		}
	}
	return
}

func WriteConfig(filename string) error {
	return os.WriteFile(
		fmt.Sprintf("%s/%s", fig.DefaultDir, filename),
		[]byte(DEFAULT_CONTENT),
		0644,
	)
}
