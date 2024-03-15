package app

import (
	"context"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func (a *App) loadRoutes(ctx context.Context) {
	router := chi.NewRouter()

	router.Use(middleware.Logger)

	router.Route("/auth", a.loadAuthRoutes)

	a.router = router
}

func (a *App) loadAuthRoutes(router chi.Router) {
	authHandler := a.serviceProvider.authAPI
	router.Post("/login", authHandler.Login)
	router.Post("/refresh", authHandler.Refresh)

}
