package database

import "github.com/ProSt1ll/wb-l0/internal/models"

type Database interface {
	Save(order models.Order) error
	Load(uid string) (models.Order, bool)
}
