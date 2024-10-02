package config

import (
	"fmt"
	"time"

	"github.com/kvizyx/glich/shared/config"
)

type Config struct {
	Shared config.Shared

	App struct {
		Mode     config.AppMode `env:"APP_MODE" env-default:"development"`
		LogLevel string         `env:"APP_LOG_LEVEL" env-default:"debug"`
	}

	HTTP struct {
		Port uint16 `env:"HTTP_PORT" env-default:"8080"`
	}

	Auth struct {
		SessionTTL time.Duration `env:"AUTH_SESSION_TTL" env-default:"30m"`
	}
}

func New(local, shared string) (Config, error) {
	builder := config.NewBuidler()

	if len(paths) == 0 {
		config, err := builder.Build()
		if err != nil {
			return Config{}, fmt.Errorf("build config: %w", err)
		}

		return config, nil
	}

	for _, path := range paths {
		builder.AddSource(config.Source{
			Kind: config.SourceKindEnv,
			Path: path,
		})
	}

	config, err := builder.Build

	return Config{}, nil
}
