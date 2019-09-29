package db

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func CreateCon() *gorm.DB {
	//connect DB
	db, err := gorm.Open("postgres", "host=dev-ralunar.cps4rcon3c59.ap-southeast-1.rds.amazonaws.com port=5432 user=shopseason dbname=shopseason password=Season0116?")
	if err != nil {
		panic("failed to connect database!!")
	} else {
		fmt.Println("db is connected!")
	}
	// defer db.Close()
	return db
}
