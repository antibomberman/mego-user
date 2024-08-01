package database

import (
	"fmt"
	"github.com/antibomberman/mego-user/internal/config"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"log"
	"time"
)

func ConnectToDB(cfg *config.Config) (*sqlx.DB, error) {
	var db *sqlx.DB
	var err error
	maxAttempts := 10
	for i := 0; i < maxAttempts; i++ {
		log.Printf("connenting to %s", cfg.DatabaseURL)
		db, err = sqlx.Open("postgres", cfg.DatabaseURL)
		if err == nil {
			err = db.Ping()
			if err == nil {
				log.Printf("Connected to database")
				return db, nil
			}
		}
		log.Printf("Failed to connect to database, retrying in 5 seconds...")
		time.Sleep(5 * time.Second)
	}
	return nil, fmt.Errorf("failed to connect to database after 10 attempts")
}
