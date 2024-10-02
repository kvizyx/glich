package http

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/kvizyx/glich/services/api/internal/config"
	"github.com/kvizyx/glich/shared/logger"
	"github.com/labstack/echo/v4"
)

type Delivery struct {
	server *echo.Echo

	config config.Config
	logger logger.StructuralLogger
}

type Options struct {
	Config config.Config
	Logger logger.StructuralLogger
}

func NewDelivery(opts Options) Delivery {
	server := echo.New()

	server.Server.IdleTimeout = opts.Config.HTTP.IdleTimeout
	server.Server.MaxHeaderBytes = opts.Config.HTTP.MaxHeaderBytes
	server.Server.ReadHeaderTimeout = opts.Config.HTTP.ReadHeaderTimeout
	server.Server.ReadTimeout = opts.Config.HTTP.ReadTimeout
	server.Server.WriteTimeout = opts.Config.HTTP.WriteTimeout

	return Delivery{
		server: server,
		config: opts.Config,
		logger: opts.Logger,
	}
}

func (d *Delivery) Start() error {
	addr := fmt.Sprintf("%s:%d", d.config.HTTP.Host, d.config.HTTP.Port)

	d.logger.Info("http server started", slog.String("addr", addr))

	if err := d.server.Start(addr); err != nil {
		if errors.Is(err, http.ErrServerClosed) {
			return nil
		}

		return fmt.Errorf("start server: %w", err)
	}

	return nil
}

func (d *Delivery) Stop(ctx context.Context) error {
	if err := d.server.Shutdown(ctx); err != nil {
		return fmt.Errorf("stop server: %w", err)
	}

	return nil
}
