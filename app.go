package main

import (
	"github.com/MySmartFarm/mysmartfarm_api/controllers"
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func main() {

	//http API
	r := gin.Default()
	r.GET("/getHello/", controllers.GetHello)

	r.Run() //listen and serve on 0.0.0.0:8080
}