package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

//(W)RAP BATTLES MENU CONFIGURATION
type Menu struct {
	gorm.Model
	menu_list []WrapItem
}

type WrapItem struct {
	gorm.Model
	Item_id     int "gorm: `unique_index`"
	Name        string
	Description string
	Recipe      string
	Price       int
	Likes       int
	// reviews     Reviews
}

// type Reviews struct {
// 	reviewer string
// 	content  []string
// date     time.Now()
// }

//Actual Menu Items for DB
var the_fkn_menu Menu = Menu{
	menu_list: []WrapItem{
		{
			Item_id:     1,
			Name:        "The Biggie",
			Description: "A Big Ass Chopped Cheese. Baybay Baybay!!",
			Recipe:      "1/2 LB Ground Beef, American Cheese, LTOP, The Sauce",
			Price:       10,
			Likes:       0,
		},
		{
			Item_id:     2,
			Name:        "The Slim Shady",
			Description: "Hi!!! My Name Is A Dope Ass Detroit Coney!",
			Recipe:      "1/4 LB All Beef Hot Dog, House Chili, Diced Onions, Mustard, Grated Cheddar('Chilllll')",
			Price:       10,
			Likes:       0,
		},
		{
			Item_id:     3,
			Name:        "The Snoop",
			Description: "Chicken+Bacon+Avocado+Ranch. Cop It Like It's Hot.",
			Recipe:      "Grilled Chicken Breast, Mashed Avocado, Bacon, Ranch, Swiss Cheese",
			Price:       10,
			Likes:       0,
		},
		{
			Item_id:     4,
			Name:        "The HNDRXX",
			Description: "Eat This Soul Food In Some Gucci Flip Flops",
			Recipe:      "Fried Porkchop, Dirty Sprite Braised Collards, Pimiento Cheese, Pinto Beans and Rice",
			Price:       10,
			Likes:       0,
		},
		{
			Item_id:     5,
			Name:        "The Abbott",
			Description: "Vegan Banh Mi's Ain't Nothin To F*ck With!",
			Recipe:      "Roasted Shiitakes and Cauliflower, Radish-Carrot Slaw, Sriracha Aioli, Mint",
			Price:       10,
			Likes:       0,
		},
		{
			Item_id:     6,
			Name:        "House Chips",
			Description: "Kettle Cooked Chips. They're good AF.",
			Recipe:      "Available in Smoked Gouda, Creamy Dill Pickle, Barbecue, Plain",
			Price:       4,
			Likes:       0,
		},
	},
}

var db *gorm.DB
var err error

func fiberHandler(c *fiber.Ctx) {
	c.Send("We outchea!")
}
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
	db.AutoMigrate(&WrapItem{})
	db.Create(&the_fkn_menu.menu_list)

	//GOFiber Instance Initialization
	app := fiber.New()
	app.Get("/", fiberHandler)
	app.Listen(":3000")
	fmt.Println("Listening on Port 3000")
}
