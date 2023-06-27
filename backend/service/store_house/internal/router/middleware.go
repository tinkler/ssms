package router

import (
	"github.com/go-chi/chi/v5"
	"github.com/tinkler/ssms/backend/service/store_house/internal/middleware"
)

func GetGlobalMiddlewares(m *chi.Mux) {
	m.Use(middleware.Cors)
}
