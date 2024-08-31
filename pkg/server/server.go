package server

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/jinzhu/gorm"
)

type Svr struct {
	Database *gorm.DB
}

// this is how we should write get api:-
func (s *Svr) HealthCheck(c *fiber.Ctx) error {
	return c.JSON(map[string]interface{}{
		"Status": "OK",
		"result": "Success",
	})
}

//this is how we should write post api:-
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

func (s *Svr) CreateTask(c *fiber.Ctx) error {

	name := c.FormValue("name")
	description := c.FormValue("description")
	status := c.FormValue("status")
	employee_id := c.FormValue("employee_id")
	project_id := c.FormValue("project_id")

	err := s.Database.Exec("INSERT INTO tasks (name, description, status, employeeid, projectid) VALUES(?,?,?,?,?)", name, description, status, employee_id, project_id).Error

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

func (s *Svr) mapping(c *fiber.Ctx) error {

	project_id := c.FormValue("project_id")
	manager_id := c.FormValue("manager_id")
	employee_id := c.FormValue("employee_id")

	err := s.Database.Exec("INSERT INTO tasks (projectid, managerid, employeeid ) VALUES(?,?,?)", project_id, manager_id, employee_id).Error

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

func (s *Svr) addEmployee(c *fiber.Ctx) error {

	name := c.FormValue("name")
	email := c.FormValue("email")
	designation := c.FormValue("designation")

	err := s.Database.Exec("INSERT INTO employees (name, email, designation) VALUES(?,?,?)", name, email, designation).Error

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
