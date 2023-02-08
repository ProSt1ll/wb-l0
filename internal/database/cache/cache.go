package cache

import (
	"github.com/ProSt1ll/wb-l0/internal/models"
)

type Cache struct {
	cache map[string]models.Order
}

func New() Cache {
	return Cache{
		cache: make(map[string]models.Order),
	}
}

func (c *Cache) Save(order models.Order) error {
	c.cache[order.OrderUID] = order
	return nil
}

func (c *Cache) Load(uid string) (models.Order, bool) {
	order := c.cache[uid]
	if len(order.OrderUID) == 0 {
		return order, false
	}
	return order, true
}
