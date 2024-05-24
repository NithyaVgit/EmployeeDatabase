package main

import (
	"fmt"
	"go/employee/attendance/config"
	mysqldb "go/employee/attendance/respoistory/Mysql"
	router "go/employee/attendance/router"
	"go/employee/attendance/validation"
)

func main() {
	configState, configError := config.ReadConfig("config/config.yml")

	if configState {
		fmt.Println("Server Start")
		mysqldb.Connection()
		validation.InitializeValidation()
		router.InitializeRouter()
	} else {
		panic(configError.Error())
	}
}
