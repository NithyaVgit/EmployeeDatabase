package respoistory

import (
	"fmt"
	"go/employee/attendance/config"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	MysqlDB *gorm.DB
)

func Connection() {

	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=true",
		config.Configs.Mysql.Username,
		config.Configs.Mysql.Password,
		config.Configs.Mysql.Host,
		config.Configs.Mysql.DatabaseName,
	)

	// Connection Instance
	database, err := gorm.Open(mysql.Open(dataSourceName), &gorm.Config{})

	if err != nil {
		panic(err.Error())
	}

	// Table Auto Migration
	migrateErr := database.AutoMigrate(&Employee{})
	if migrateErr != nil {
		panic(migrateErr.Error())
	}

	MysqlDB = database
	log.Println("Mysql Connection Successfully")

}
