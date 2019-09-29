package main

import (
	"api_get_products/controllers"
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func main() {

	//http API
	r := gin.Default()
	r.GET("/getProduct/:id", controllers.GetProductFunc)

	r.Run() //listen and serve on 0.0.0.0:8080
}