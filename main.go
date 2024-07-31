package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
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
