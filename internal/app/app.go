package app

import (
	"github.com/ProSt1ll/wb-l0/internal/database"
	"github.com/ProSt1ll/wb-l0/internal/transport/rest"
)

type App struct {
	server rest.Rest
	db     database.Database
}

func New() App {
	db := database.New()
	return App{
		db:     &db,
		server: rest.New(&db),
	}
}

func (a *App) Run() error {
	return a.server.Run()
}
