package database

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/tooth-fairy/config"
)

// Database - Struct that represents a database connection
type Database struct {
	client *gorm.DB
}

//Close ....
func (conn *Database) Close() error {
	return conn.client.Close()
}

func createConnection(config *config.AppConfig) (*gorm.DB, error) {
	DBURL := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s",
		config.DatabaseHost, config.DatabasePort, config.DatabaseUser, config.DatabaseName, config.DatabasePass)
	db, err := gorm.Open("postgres", DBURL)
	if err != nil {
		return nil, err
	}

	return db, nil
}

//CheckLiveness - verify if database is alive
func (conn *Database) CheckLiveness() error {
	_db := conn.client.DB()
	err := _db.Ping()
	if err != nil {
		_db.Close()
	}
	return nil
}

//New - Creates a new database object
func New(config *config.AppConfig) (*Database, error) {
	client, err := createConnection(config)

	if err != nil {
		return nil, err
	}

	return &Database{
		client: client,
	}, nil
}
