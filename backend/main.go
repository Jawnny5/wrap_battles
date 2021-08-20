package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/postgres"
	_ "gorm.io/driver/postgres"
	"gorm.io/gorm"
)

//(W)RAP BATTLES MENU CONFIGURATION
type Menu struct {
	gorm.Model

	item_id     int
	name        string
	description string
	price       int
	tags        []string
	likes       int
	reviews     string
}

var db *gorm.DB
var err error

func main() {
	//IMPORTING ENV VARIABLES
	// dialect := os.Getenv("DIALECT")
	host := os.Getenv("HOST")
	dbport := os.Getenv("DBPORT")
	dbname := os.Getenv("NAME")
	dbuser := os.Getenv("USER")

	//DB CONNECTION STRING
	dbURI := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable port=%s", host, dbuser, dbname, dbport)

	//OPEN CONNECTION TO DATABASE
	db, err = gorm.Open(postgres.Open(dbURI), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("Successfully connected to RAPBTLZ Menu DB")
	}

	//Close connection to database when main function is finished
	menu_db, _ := db.DB()
	defer menu_db.Close()

	//Configuration of PSQL Migrations in GORM
	db.AutoMigrate(&Menu{})

	//GOFiber Instance Initialization
	app := fiber.New()
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("We outchea")
	})
	app.Listen(":3000")
	fmt.Println("Listening on Port 3000")
}
