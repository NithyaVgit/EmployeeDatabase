package router

import (
	"go/employee/attendance/config"
	"go/employee/attendance/handler"
	"net"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func employeeinsertRouters(routers *gin.RouterGroup) {
	routers.POST("/addemployee", func(c *gin.Context) {
		handler.EmployeeInsertHandler(c.Writer, c.Request)
	})
}

func getemployeeRouters(routers *gin.RouterGroup) {
	routers.GET("/getemployee", func(c *gin.Context) {
		handler.GetEmployeeHandler(c.Writer, c.Request)
	})
}

func getallemployeeRouters(routers *gin.RouterGroup) {
	routers.GET("/getallemployee", func(c *gin.Context) {
		handler.GetAllEmployeeHandler(c.Writer, c.Request)
	})
}

func employeeupdateRouters(routers *gin.RouterGroup) {
	routers.PATCH("/updateemployee", func(c *gin.Context) {
		handler.EmployeeUpdateHandler(c.Writer, c.Request)
	})
}

func employeeinsertcsvRouters(routers *gin.RouterGroup) {
	routers.POST("/addcsveemployee", func(c *gin.Context) {
		handler.EmployeecsvInsertHandler(c.Writer, c.Request)
	})
}

func employeedownloadcsvRouters(routers *gin.RouterGroup) {
	routers.POST("/downloadcsvemployee", func(c *gin.Context) {
		handler.EmployeeCsvDownloadHandler(c.Writer, c.Request)
	})
}

func InitializeRouter() {

	gin.SetMode(gin.ReleaseMode)
	router := gin.New()                // Create Gin Router
	routers := router.Group("/api/v1") // Create Router Group
	employeeinsertRouters(routers)
	getemployeeRouters(routers)
	getallemployeeRouters(routers)
	employeeupdateRouters(routers)
	employeeinsertcsvRouters(routers)
	employeedownloadcsvRouters(routers)

	/*
			Set the read and write timeouts for the HTTP server
			The address to listen on
			The Gin router to handle requests
			Maximum time to read the request from the client
		    Maximum time to write the response back to the client
	*/
	server := &http.Server{
		Addr:         config.Configs.App.Host + ":" + config.Configs.App.Port,
		Handler:      router,
		ReadTimeout:  120 * time.Second,
		WriteTimeout: 120 * time.Second,
	}

	listener, err := net.Listen("tcp4", config.Configs.App.Host+":"+config.Configs.App.Port)
	if err != nil {
		panic(err.Error())
	}
	serverErr := server.Serve(listener)

	if serverErr != nil && serverErr != http.ErrServerClosed {
		panic(serverErr.Error())
	}

}
