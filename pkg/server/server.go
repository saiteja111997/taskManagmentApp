package server

import (
	"time"

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

// func (s *Svr) Demo(c *fiber.Ctx) error {
// 	message := c.FormValue("message")
// 	fmt.Println("Message Content : ", message)

// 	if message == "" {
// 		return c.JSON(map[string]interface{}{
// 			"Status": "OK",
// 			"result": "empty message",
// 		})
// 	} else {
// 		return c.JSON(map[string]interface{}{
// 			"Status": "OK",
// 			"result": message,
// 		})
// 	}

// }

func (s *Svr) CreateProject(c *fiber.Ctx) error {

	name := c.FormValue("name")
	description := c.FormValue("description")
	category := c.FormValue("category")
	status := c.FormValue("status")

	startDate := time.Now()

	err := s.Database.Exec("INSERT INTO project_infos (name, description, category, status, start_date) VALUES(?,?,?,?,?)", name, description, category, status, startDate).Error

	if err != nil {
		return c.JSON(map[string]interface{}{
			"Status": "!OK",
			"result": err,
		})
	}

	return c.JSON(map[string]interface{}{
		"Status": "OK",
		"result": "Successfully inserted into the database",
	})
}
