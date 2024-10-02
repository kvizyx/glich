package config

import (
	"fmt"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
	"github.com/joho/godotenv"
	"github.com/kvizyx/glich/shared/config"
)

type Config struct {
	Shared config.Shared

	App struct {
		Mode     config.AppMode `env:"APP_MODE" env-default:"development"`
		LogLevel string         `env:"APP_LOG_LEVEL" env-default:"debug"`
	}

	HTTP struct {
		Host string `env:"HTTP_HOST" env-default:"127.0.0.1"`
		Port uint16 `env:"HTTP_PORT" env-default:"8080"`

		IdleTimeout       time.Duration `env:"HTTP_IDLE_TIMEOUT" env-default:""`
		WriteTimeout      time.Duration `env:"HTTP_WRITE_TIMEOUT" env-default:""`
		ReadTimeout       time.Duration `env:"HTTP_READ_TIMEOUT" env-default:""`
		ReadHeaderTimeout time.Duration `env:"HTTP_READ_HEADER_TIMEOUT" env-default:""`
		MaxHeaderBytes    int           `env:"HTTP_MAX_HEADER_BYTES" env-default:""`
	}

	Auth struct {
		SessionTTL time.Duration `env:"AUTH_SESSION_TTL" env-default:"30m"`
	}
}

func New(local, shared string) (Config, error) {
	var (
		config Config
		scopes = [2]string{local, shared}
	)

	for _, scope := range scopes {
		if len(scope) == 0 {
			continue
		}

		if err := godotenv.Load(scope); err != nil {
			return config, fmt.Errorf("load scope environment variables: %w", err)
		}
	}

	if err := cleanenv.ReadEnv(&config); err != nil {
		return config, fmt.Errorf("read environment variables into config: %w", err)
	}

	return config, nil
}
