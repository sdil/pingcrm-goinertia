package server

import (
	"database/sql"
)

type Container struct {
	DB *sql.DB
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
	err := c.DB.Close()
	if err != nil {
		return err
	}

	return nil
}
