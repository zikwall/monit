package middlewares

import "github.com/gofiber/fiber/v2"

func IP(c *fiber.Ctx) error {
	ip := c.IP()

	if c.Get("X-Real-IP") != "" {
		ip = c.Get("X-Real-IP")
	} else if c.Get("X-Forwarded-For") != "" {
		if ips := c.IPs(); len(ips) > 0 {
			ip = ips[0]
		}
	}

	c.Locals("ip", ip)
	return c.Next()
}
