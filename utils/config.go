package utils

import (
	"os"
	"strings"

	"github.com/disgoorg/log"
	"github.com/disgoorg/snowflake/v2"
	"github.com/spf13/viper"
)

const ENVVAR_PREFIX = "env:"

func SetupConfig(l log.Logger, path string) *viper.Viper {
	v := viper.NewWithOptions(
		viper.KeyDelimiter("~"),
	)
	v.SetConfigName("botconfig")
	v.SetConfigType("yaml")
	v.AddConfigPath(path)

	v.SetDefault("debug-log-level", false)
	v.SetDefault("bot~token", "")
	v.SetDefault("bot~setup-commands", true)
	v.SetDefault("bot~global-commands", false)
	v.SetDefault("bot~guild-id", snowflake.ID(0))
	v.SetDefault("bot~mobile-os", true)
	// v.SetDefault("", "")

	if err := v.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			l.Infof("Config file not found, creating in '%s' path", path)
			if err := v.WriteConfigAs("botconfig.yml"); err != nil {
				l.Fatal("Config file writer error: ", err)
			}
		} else {
			l.Fatal("Config file reader error: ", err)
		}
	}

	for _, key := range v.AllKeys() {
		if value := v.GetString(key); value != "" && strings.HasPrefix(
			value,
			ENVVAR_PREFIX,
		) {
			v.Set(
				key,
				os.Getenv(value[len(ENVVAR_PREFIX):]),
			)
		}
	}

	return v
}
