package contract

import "github.com/gofiber/fiber/v2"

// Container represents a contract for a container
type Container interface {
	HandleApi(r fiber.Router)
	HandleWeb(r fiber.Router)
}
