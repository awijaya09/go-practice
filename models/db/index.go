package db

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var DB *gorm.DB
var err error

// Connect to DB in MySQL
func Connect() {
	DB, err = gorm.Open("postgres", "postgres://sdwmmmvjwurjru:85a060bbe732512fff3b16723bb47b2e3d24ffd236f7262e78c311f313313c0a@ec2-50-17-227-28.compute-1.amazonaws.com:5432/d5ukif75k3i535")

	if err != nil {
		fmt.Println("Unable to connect to DB...")
		fmt.Println("Error >> :", err.Error())
		DB.Close()
		return
	}

	migrate(DB)
}

func migrate(db *gorm.DB) {
	db.AutoMigrate(&User{})
}
