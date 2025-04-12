package app

import (
	"database/sql"
	"fmt"
	"os"
	"time"

	"github.com/abdultalif/restful-api/helper"
	"github.com/joho/godotenv"
)

func NewDB() *sql.DB {

	err := godotenv.Load()
	helper.PanicIfError(err)

	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbDialect := os.Getenv("DB_DIALECT")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", 
		dbUser, dbPassword, dbHost, dbPort, dbName)

	db, err := sql.Open(dbDialect, dsn)
	helper.PanicIfError(err)

	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(20)
	db.SetConnMaxLifetime(60 * time.Minute)
	db.SetConnMaxIdleTime(10 * time.Minute)

	return db
}