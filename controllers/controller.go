package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetHello(c *gin.Context) {
	//Parameter from url
	fmt.Println("Hello")
	//Return json
	c.JSON(http.StatusOK, "")
}
