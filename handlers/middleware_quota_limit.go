package handlers

import (
	"sync"

	"github.com/gofiber/fiber/v2"
)

var ipQuota = struct {
	sync.Map
}{}

const defaultQuota = 100

func QuotaLimit(c *fiber.Ctx) error {
	ip := c.IP()

	val, ok := ipQuota.Load(ip)
	if !ok {
		ipQuota.Store(ip, defaultQuota)
		val = defaultQuota
	}

	quota := val.(int)
	if quota <= 0 {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"message": "API quota exceeded for your IP",
		})
	}
	ipQuota.Store(ip, quota-1)
	return c.Next()
}
