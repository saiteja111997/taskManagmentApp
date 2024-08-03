package database

import (
	"fmt"
	"taskManagmentApp/pkg/structures"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func ConnectToDatabase(config structures.DbConfig) *gorm.DB {
	connectionString := ConnectString(config)

	db, err := gorm.Open("postgres", connectionString)
	if err != nil {
		panic(err)
	}
	db.LogMode(false)
	return db

}

func ConnectString(config structures.DbConfig) string {
	var str string

	str = fmt.Sprintf(`host=%v port=%v user=%v dbname=%v password=%v sslmode=disable`,
		config.DB_HOSTNAME,
		config.DB_PORT,
		config.DB_USERNAME,
		config.DB_NAME,
		config.DB_PASSWORD,
	)

	return str
}
