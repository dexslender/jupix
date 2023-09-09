package util

import (
	"errors"
	"fmt"
	"os"

	"github.com/disgoorg/log"
	"github.com/disgoorg/snowflake/v2"
	"github.com/kkyr/fig"
)

var DEFAULT_CONTENT = `
	bot:
		token: your-token-here
		setup-commands: false
		global-commands: false
		guild-id: 0
		mobile-os: false
	config:
		log-debug: false
`

type Config struct {
	Bot struct {
		Token          string       `validate:"required"`
		SetupCommands  bool         `fig:"setup-commands" default:"false"`
		GlobalCommands bool         `fig:"global-commands" default:"false"`
		GuildId        snowflake.ID `fig:"guild-id"`
		MobileOs       bool         `fig:"mobile-os" default:"false"`
	}
	Config struct {
		LogDebug bool `fig:"log-debug" default:"false"`
	}
}

func LoadConfig(l log.Logger, filename string) (cfg Config) {
	if err := fig.Load(
		&cfg,
		fig.File(filename),
		// fig.UseEnv("PLANE"),
	); err != nil {
		if errors.Is(err, error(fig.ErrFileNotFound)) {
			l.Infof("Config file not found!")
			l.Infof("Config file not found creating in '%s'", fig.DefaultDir)

			if err := WriteConfig(filename); err != nil {
				log.Errorf("Error while writting config file: ", err)
			}
		} else {
			l.Fatal("Config reader error: ", err)
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
