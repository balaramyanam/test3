package main

import (
	//"database/sql"
	_ "encoding/json"
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/gofiber/fiber/v2"
	_ "github.com/jinzhu/gorm"
)

type Phones struct {
	Name string
	Cost int
}

var (
	DBConn *gorm.DB
)

func main() {

	dbn := "root:Balaram@123@tcp(127.0.0.1:3306)/sys?charset=utf8&parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.Open(dbn), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect databse. \n", err)
		os.Exit(2)
	}
	log.Println("connected")
	db.AutoMigrate(&Phones{})
	DBConn = db
	bhavana := Phones{"realme", 25000}

	ap := fiber.New()
	ap.Post("/smiling", func(c *fiber.Ctx) error {
		DBConn.Create(&bhavana)

		return c.Status(404).JSON(&bhavana)

	})
	log.Fatal(ap.Listen(":2003"))
}
