package api

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

// SetRoutes sets the routes with
// corresponding handler functions
func (s *Server) SetRoutes() {
	s.r.Use(middleware.Logger)
	s.currencyRoutes()
	s.currencyRateRoutes()
	s.fanRoutes()
	s.r.Post("/create/creator", s.CreateCreator(s.ctx))
}

// Sets up routes for currency
func (s *Server) currencyRoutes() {
	s.r.Route("/currency", func(r chi.Router) {
		r.Get("/", s.ListCurrencies(s.ctx))
		r.Post("/", s.CreateCurrency(s.ctx))

		// Subroutes for currency
		r.Route("/{id}", func(r chi.Router) {
			r.Get("/", s.GetCurrency(s.ctx))
			r.Put("/", s.UpdateCurrency(s.ctx))
			r.Delete("/", s.DeleteCurrency(s.ctx))
		})
	})
}

// Sets up routes for currency rates
func (s *Server) currencyRateRoutes() {
	s.r.Route("/currency/rate", func(r chi.Router) {
		r.Post("/", s.CreateCurrencyRate(s.ctx))

		r.Route("/{id}", func(r chi.Router) {
			r.Get("/", s.GetCurrencyRate(s.ctx))
			r.Put("/", s.UpdateCurrencyRate(s.ctx))
		})
	})
}

// Sets up routes for fan
func (s *Server) fanRoutes() {
	s.r.Route("/fan", func(r chi.Router) {
		r.Post("/", s.CreateFan(s.ctx))

		r.Route("/{id}", func(r chi.Router) {
			r.Get("/", s.GetFan(s.ctx))
			r.Put("/email", s.UpdateFanEmail(s.ctx))
			r.Put("/password", s.UpdateFanPassword(s.ctx))
			r.Put("/pcurrency", s.UpdateFanPreferredCurrency(s.ctx))
		})
	})
}
