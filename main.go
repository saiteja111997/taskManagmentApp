package main

import (
	"context"
	"fmt"
	database "taskManagmentApp/pkg/db"
	"taskManagmentApp/pkg/server"
	"taskManagmentApp/pkg/structures"
	"taskManagmentApp/pkg/utilities"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	fiberadapter "github.com/awslabs/aws-lambda-go-api-proxy/fiber"

	"github.com/gofiber/fiber/v2"
	"github.com/jinzhu/gorm"
)

var fiberLambda *fiberadapter.FiberLambda

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

	db.AutoMigrate(&structures.ProjectInfo{}, &structures.Mapping{}, &structures.Employee{}, &structures.Task{})

	svr := server.Svr{
		Database: db,
	}

	// var svr server.Svr
	// svr.Database = db

	// HEALTH CHECK OR HEART BEAT
	app.Get("/healthCheck", svr.HealthCheck)
	//for creating tables and adding data into it
	app.Post("/createProject", svr.CreateProject)
	app.Post("/createTask", svr.CreateTask)
	app.Post("/mapping", svr.CreateMapping)
	app.Post("/addEmployee", svr.AddEmployee)
	app.Post("/getTaskStatus", svr.TaskStatus)
	app.Post("/updateTaskStatus", svr.UpdateTaskStatus)
	app.Get("/getEmployees", svr.GetEmployees)

	// Payload types : JSON, XML, FORMDATA etc...
	// app.Post("/demo", svr.Demo)

	// frontend => data (JSON) ([]byte)
	//Backend => Unmarshal([]byte -> JSON) (Read keys)

	if utilities.IsLambda() {
		fiberLambda = fiberadapter.New(app)
		lambda.Start(Handler)
	} else {
		fmt.Println("Starting server locally!!")
		err = app.Listen(":8090")

		if err != nil {
			fmt.Println("An error occured while starting the server : ", err)
		}
	}
}

func Handler(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	// Proxy the request to the Fiber app and get the response
	response, err := fiberLambda.ProxyWithContext(ctx, request)

	response.Headers = make(map[string]string)

	// Add CORS headers to the response
	response.Headers["Access-Control-Allow-Origin"] = "*"
	response.Headers["Access-Control-Allow-Methods"] = "GET,POST,PUT,DELETE"
	response.Headers["Access-Control-Allow-Headers"] = "Origin, Content-Type, Accept"
	response.Headers["Access-Control-Allow-Credentials"] = "true"

	return response, err
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
