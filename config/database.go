package config

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"pioApi/ent"
	"time"

	"entgo.io/ent/dialect"
	entsql "entgo.io/ent/dialect/sql"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/joho/godotenv"
)

var DB *ent.Client

func ConnectDB() error {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file: ", err)
	}

	dbStr := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
	)

	db, err := sql.Open("pgx", dbStr)
	if err != nil {
		log.Fatal("Failed to connect to the database: ", err)
		return err
	}

	db.SetMaxOpenConns(20)
	db.SetMaxIdleConns(10)
	db.SetConnMaxLifetime(30 * time.Minute)

	drv := entsql.OpenDB(dialect.Postgres, db)
	fmt.Println("Database connected with connection pooling enabled!")

	DB = ent.NewClient(ent.Driver(drv))
	return nil
}
