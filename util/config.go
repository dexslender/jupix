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
  guild: 0
  setup-commands: false
  global-commands: false
  use-mobile-os: true
  presence-interval: 20s

runtime:
  debug-log: false

presences:
  - name: discord
    status: online

  # - name: 
  #   type:
  #   status:
  #   url:
  #	  state:
`

type Config struct {
	Bot struct {
		Token            string        `validate:"required"`
		GuildId          snowflake.ID  `fig:"guild" default:"0"`
		SetupCommands    bool          `fig:"setup-commands"`
		GlobalCommands   bool          `fig:"global-commands"`
		MobileOs         bool          `fig:"use-mobile-os"`
		PresenceInterval time.Duration `fig:"presence-interval" default:"10s"`
	}
	Runtime struct {
		DebugLog bool `fig:"debug-log"`
	}
	Presences []Presence
}

type Presence struct {
	Status discord.OnlineStatus
	Name   string
	Type   string
	URL    string
	State  string
}

func LoadConfig(l log.Logger, filename string) (config Config) {
	err := fig.Load(
		&config,
		fig.UseEnv("JUPIX"),
		fig.File(filename),
	)
	if err != nil {
		if errors.Is(err, fig.ErrFileNotFound) {
			l.Infof("Config file not found, creating in '%s'", fig.DefaultDir)
			if err := WriteConfig(filename); err != nil {
				log.Error("Error while written config file: ", err)
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

// package util

// import (
// 	"os"
// 	"strings"

// 	"github.com/disgoorg/log"
// 	"github.com/disgoorg/snowflake/v2"
// 	"github.com/spf13/viper"
// )

// const ENVVAR_PREFIX = "env:"

// func SetupConfig(l log.Logger, path string) *viper.Viper {
// 	v := viper.NewWithOptions(
// 		viper.KeyDelimiter("~"),
// 	)
// 	v.SetConfigName("botconfig")
// 	v.SetConfigType("yaml")
// 	v.AddConfigPath(path)

// 	v.SetDefault("debug-log-level", false)
// 	v.SetDefault("bot~token", "")
// 	v.SetDefault("bot~setup-commands", true)
// 	v.SetDefault("bot~global-commands", false)
// 	v.SetDefault("bot~guild-id", snowflake.ID(0))
// 	v.SetDefault("bot~mobile-os", true)
// 	// v.SetDefault("", "")

// 	if err := v.ReadInConfig(); err != nil {
// 		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
// 			l.Infof("Config file not found, creating in '%s' path", path)
// 			if err := v.WriteConfigAs("botconfig.yml"); err != nil {
// 				l.Fatal("Config file writer error: ", err)
// 			}
// 		} else {
// 			l.Fatal("Config file reader error: ", err)
// 		}
// 	}

// 	for _, key := range v.AllKeys() {
// 		if value := v.GetString(key); value != "" && strings.HasPrefix(
// 			value,
// 			ENVVAR_PREFIX,
// 		) {
// 			v.Set(
// 				key,
// 				os.Getenv(value[len(ENVVAR_PREFIX):]),
// 			)
// 		}
// 	}

// 	return v
// }
