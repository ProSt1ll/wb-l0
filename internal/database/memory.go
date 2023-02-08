package database

import (
	"fmt"
	"github.com/ProSt1ll/wb-l0/internal/database/cache"
	"github.com/ProSt1ll/wb-l0/internal/database/sql"
	"github.com/ProSt1ll/wb-l0/internal/models"
	"log"
)

type Memory struct {
	Sql   sql.SQL
	Cache cache.Cache
}

func New() Memory {
	m := Memory{
		Sql:   sql.New(),
		Cache: cache.New(),
	}
	m.DBtoMem()
	return m
}

func (m *Memory) DBtoMem() {
	orders, ok := m.Sql.LoadAll()
	if !ok {
		return
	}
	for _, order := range orders {
		if err := m.Cache.Save(order); err != nil {
			log.Fatal(err)
		}
	}
	return
}

func (m *Memory) Save(order models.Order) error {
	fmt.Println("saved id : " + order.OrderUID)
	if err := m.Cache.Save(order); err != nil {
		return err
	}
	if err := m.Sql.Save(order); err != nil {
		return err
	}
	return nil
}

func (m *Memory) Load(uid string) (models.Order, bool) {
	return m.Cache.Load(uid)
}
