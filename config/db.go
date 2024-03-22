package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDB() *gorm.DB {
	var err error

	host := os.Getenv("DBHOST")
	port, err := strconv.Atoi(os.Getenv("DBPORT"))
	user := os.Getenv("DBUSER")
	password := os.Getenv("DBPASSWORD")
	dbname := os.Getenv("DBNAME")

	if err != nil {
		log.Println(err.Error())
	}

	// Generate Connection String
	connTmpl := "host=%s port=%d dbname=%s user=%s password=%s sslmode=disable"

	// Another template
	// connTmpl := "postgresql://%s:%s@%s:%d/%s?sslmode=disable"
	//user:password@host:port/dbname

	connStr := fmt.Sprintf(
		connTmpl,
		host,
		port,
		dbname,
		user,
		password,
	)

	// Connect to database
	db, err := gorm.Open(postgres.Open(connStr), &gorm.Config{})
	if err != nil {
		log.Println("Cannot connect to database")
	}

	return db
}
