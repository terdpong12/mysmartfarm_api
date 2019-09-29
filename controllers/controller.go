package controllers

import (
	"github.com/MySmartFarm/mysmartfarm_api/models"
	"fmt"
	"net/http"
	"github.com/gin-gonic/gin"
)

func GetProductFunc(c *gin.Context) {
	//Parameter from url
	productID := c.Param("id")
	result := models.GetProd1(productID)
	fmt.Println(productID)
	fmt.Println(result)
	//Return json
	c.JSON(http.StatusOK, gin.H{"Products": result})
}

func GetHello(c *gin.Context) {
	//Parameter from url
	fmt.Println("Hello")
	//Return json
	c.JSON(http.StatusOK, "")
}
