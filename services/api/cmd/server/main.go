package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/kvizyx/glich/services/api/internal/config"
	"github.com/kvizyx/glich/shared/logger"
)

var (
	envLocal  string
	envShared string
)

func init() {
	flag.StringVar(&envLocal, "env-local", "", "path to the local envitonment variables")
	flag.StringVar(&envShared, "env-shared", "", "path to the shared envitonment variables")

	flag.Parse()
}

func main() {
	cfg, err := config.New(envLocal, envShared)
	if err != nil {
		fmt.Fprintf(os.Stderr, "config: %s\n", err)
		return
	}

	slg, err := logger.NewSlogLogger(logger.SlogOptions{
		AppMode: cfg.App.Mode,
		Service: "api",
		Level:   cfg.App.LogLevel,
	})
	if err != nil {
		fmt.Fprintf(os.Stderr, "logger: %s\n", err)
		return
	}

	_ = slg
}
