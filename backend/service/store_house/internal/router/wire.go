//go:build wireinject
// +build wireinject

package router

import (
	"net/http"

	"github.com/google/wire"
	"github.com/tinkler/mqttadmin/pkg/conf"
	"github.com/tinkler/mqttadmin/pkg/db"
)

var (
	confSet = wire.NewSet(conf.NewGormConfig, conf.NewConf)
	dbSet   = wire.NewSet(db.NewDB)
)

func NewServer() (*http.Server, error) {
	wire.Build(confSet, dbSet, NewRouterServer)
	return nil, nil
}
