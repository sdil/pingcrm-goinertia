package server

import (
	"gorm.io/gorm"
)

type Container struct {
	DB *gorm.DB
}

func NewContainer() *Container {
	c := new(Container)
	c.initDB()
	return c
}

func (c *Container) initDB() {
	c.DB = ConnectDb()
}

func (c *Container) Shutdown() error {
	// FIX ME: This is not the correct way to close the database connection
	a, _ := c.DB.DB()
	if err := a.Close(); err != nil {
		return err
	}

	return nil
}