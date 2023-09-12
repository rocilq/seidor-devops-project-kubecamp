package middleware

import (
	"authService/lib"
	"sync"

	"github.com/gofiber/fiber/v2"
)

var rateLimitInfo = make(lib.RateLimiterRecord)
var locker sync.Mutex

func ValidateRateLimit(c *fiber.Ctx) error {
	ip := c.IP()
	locker.Lock()
	defer locker.Unlock()
	if err := rateLimitInfo.ValidateIPRequest(ip); err != nil {
		return c.SendStatus(429)
	}
	return c.Next()
}
