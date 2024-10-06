package server

import (
	"vuego/organizations"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)
  
func ConnectDb() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&organizations.Organization{})

	return db
}