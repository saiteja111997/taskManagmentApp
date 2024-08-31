package server

import (
	"strconv"
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
	userId := c.FormValue("user_id")

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

	employee_id_int, err := strconv.Atoi(employee_id)
	if err != nil {
		return c.JSON(map[string]interface{}{
			"Status": "!OK",
			"result": err,
		})
	}

	project_id_int, err := strconv.Atoi(project_id)
	if err != nil {
		return c.JSON(map[string]interface{}{
			"Status": "!OK",
			"result": err,
		})
	}

	err = s.Database.Exec("INSERT INTO tasks (name, description, status, employee_id, project_id) VALUES(?,?,?,?,?)", name, description, status, employee_id_int, project_id_int).Error

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

func (s *Svr) CreateMapping(c *fiber.Ctx) error {

	project_id := c.FormValue("project_id")
	manager_id := c.FormValue("manager_id")
	employee_id := c.FormValue("employee_id")

	project_id_int, err := strconv.Atoi(project_id)
	if err != nil {
		return c.JSON(map[string]interface{}{
			"Status": "!OK",
			"result": err,
		})
	}

	employee_id_int, err := strconv.Atoi(employee_id)
	if err != nil {
		return c.JSON(map[string]interface{}{
			"Status": "!OK",
			"result": err,
		})
	}

	manager_id_int, err := strconv.Atoi(manager_id)
	if err != nil {
		return c.JSON(map[string]interface{}{
			"Status": "!OK",
			"result": err,
		})
	}

	err = s.Database.Exec("INSERT INTO mappings (project_id, manager_id, employee_id ) VALUES(?,?,?)", project_id_int, manager_id_int, employee_id_int).Error

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

func (s *Svr) AddEmployee(c *fiber.Ctx) error {

	name := c.FormValue("name")
	email := c.FormValue("email")
	designation := c.FormValue("designation")

	designation_int, err := strconv.Atoi(designation)
	if err != nil {
		return c.JSON(map[string]interface{}{
			"Status": "!OK",
			"result": err,
		})
	}

	err = s.Database.Exec("INSERT INTO employees (name, email, designation) VALUES(?,?,?)", name, email, designation_int).Error

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
