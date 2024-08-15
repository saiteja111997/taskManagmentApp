package server

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/jinzhu/gorm"
)

type Svr struct {
	Database *gorm.DB
}

func (s *Svr) HealthCheck(c *fiber.Ctx) error {
	return c.JSON(map[string]interface{}{
		"Status": "OK",
		"result": "Success",
	})
}

func (s *Svr) Demo(c *fiber.Ctx) error {
	message := c.FormValue("message")
	fmt.Println("Message Content : ", message)

	if message == "" {
		return c.JSON(map[string]interface{}{
			"Status": "OK",
			"result": "empty message",
		})
	} else {
		return c.JSON(map[string]interface{}{
			"Status": "OK",
			"result": message,
		})
	}

}
