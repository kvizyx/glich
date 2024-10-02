package config

import "errors"

var (
	ErrInvalidAppMode = errors.New("invalid application mode")
)

type AppMode string

const (
	AppModeDevelopment AppMode = "development"
	AppModeProduction  AppMode = "production"
)

type Shared struct {
	Postgres Postgres
}

type Postgres struct {
	URL string `env:"POSTGRES_URL" env-required:"true"`
}
