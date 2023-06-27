package router

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/tinkler/mqttadmin/pkg/conf"
	"github.com/tinkler/mqttadmin/pkg/db"
	"gorm.io/gorm"
)

func GetRoutes(m *chi.Mux) {
	// m.Route("/mqtt", func(r chi.Router) {
	// 	mqttadmin.RoutesUser(r)
	// })
}

func NewRouterServer(conf *conf.Conf, d *gorm.DB) (*http.Server, error) {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(db.WrapGorm())
	GetGlobalMiddlewares(r)
	GetRoutes(r)

	server := http.Server{
		Addr:         conf.Server.Addr,
		Handler:      r,
		ReadTimeout:  conf.Server.ReadTimeout,
		WriteTimeout: conf.Server.WriteTimeout,
		IdleTimeout:  conf.Server.IdleTimeout,
	}
	return &server, nil
}
