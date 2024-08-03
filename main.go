package main

import (
	"fmt"
	database "taskManagmentApp/pkg/db"
	"taskManagmentApp/pkg/structures"

	"github.com/gofiber/fiber/v2"
	"github.com/jinzhu/gorm"
)

func main() {
	fmt.Println("Application server started!!")

	app := fiber.New()
	var err error

	//WAITING FOR THE HOST

	// if err := waitForHost("mydbinstance.c1cnaivzlk0f.us-east-1.rds.amazonaws.com", "5432"); err != nil {
	// 	log.Fatalln(err)
	// }

	// CONNECTING TO THE DATABASE

	var db *gorm.DB

	var databaseCreds structures.DbConfig

	databaseCreds.DB_USERNAME = "postgres"
	databaseCreds.DB_PASSWORD = "7396569423"
	databaseCreds.DB_HOSTNAME = "127.0.0.1"
	databaseCreds.DB_PORT = "5432"
	databaseCreds.DB_NAME = "postgres"

	db = database.ConnectToDatabase(databaseCreds)

	err = db.DB().Ping()

	if err != nil {
		fmt.Println("Failed to ping the datbase : ", err)
	} else {
		fmt.Println("Successfully connected to the database!!")
	}

	db.AutoMigrate(&structures.ProjectInfo{})

	fmt.Println("Starting server locally!!")
	err = app.Listen(":8090")

	if err != nil {
		fmt.Println("An error occured while starting the server : ", err)
	}

}

// func waitForHost(host, port string) error {
// 	timeOut := time.Second

// 	if host == "" {
// 		return errors.Errorf("unable to connect to %v:%v", host, port)
// 	}

// 	for i := 0; i < 60; i++ {
// 		fmt.Printf("waiting for %v:%v ...\n", host, port)
// 		conn, err := net.DialTimeout("tcp", host+":"+port, timeOut)
// 		if err == nil {
// 			fmt.Println("done!")
// 			conn.Close()
// 			return nil
// 		}

// 		time.Sleep(time.Second)
// 	}

// 	return errors.Errorf("timeout attempting to connect to %v:%v", host, port)
// }
