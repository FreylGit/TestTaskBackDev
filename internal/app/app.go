package app

import (
	"context"
	"fmt"
	"github.com/FreylGit/TestTaskBackDev/internal/config"
	"net/http"
	"time"
)

type App struct {
	serviceProvider *serviceProvider
	router          http.Handler
}

func NewApp(ctx context.Context) (*App, error) {
	a := &App{}

	err := a.initDeps(ctx)
	a.loadRoutes(ctx)
	if err != nil {
		return nil, err
	}

	return a, nil
}

func (a *App) initDeps(ctx context.Context) error {
	inits := []func(ctx2 context.Context) error{
		a.initConfig,
		a.initServiceProvider,
		a.initAuthAPI,
	}

	for _, f := range inits {
		err := f(ctx)
		if err != nil {
			return err
		}
	}
	a.loadRoutes(ctx)
	return nil
}

func (a *App) initConfig(ctx context.Context) error {
	err := config.Load(".env")
	if err != nil {
		return err
	}

	return nil
}

func (a *App) initServiceProvider(ctx context.Context) error {
	a.serviceProvider = newServiceProvider()

	return nil
}

func (a *App) initAuthAPI(ctx context.Context) error {
	a.serviceProvider.AuthAPI(ctx)

	return nil
}

func (a *App) Start(ctx context.Context) error {
	server := &http.Server{
		Addr:    fmt.Sprintf(":%s", a.serviceProvider.ServerConfig().HttpPort()),
		Handler: a.router,
	}

	ch := make(chan error, 1)

	go func() {
		err := server.ListenAndServe()
		if err != nil {
			ch <- fmt.Errorf("failed to start server: %w", err)
		}
		close(ch)
	}()

	select {
	case err := <-ch:
		return err
	case <-ctx.Done():
		timeout, cancel := context.WithTimeout(context.Background(), time.Second*10)
		defer cancel()

		return server.Shutdown(timeout)
	}

	return nil
}
