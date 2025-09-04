package handlers

import (
	middleware "github.com/NiltonSousa/goapi/internal/middleware"
	"github.com/go-chi/chi"
	chimiddleware "github.com/go-chi/chi/middleware"
)

func Handler(r *chi.Mux) {
	// Defining routers
	r.Use(chimiddleware.StripSlashes)
	r.Route("/account", func(router chi.Router) {
		router.Use(middleware.Authorization)
		router.Get("/coins", GetCoinBalance)
	})
}
