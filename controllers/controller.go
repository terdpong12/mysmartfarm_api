package controllers

import (
	"api_get_products/models"
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
