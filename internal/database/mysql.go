package database

import (
	_ "github.com/go-sql-driver/mysql"
)

//func ConnectToDB(cfg *config.Config) (*sqlx.DB, error) {
//	var db *sqlx.DB
//	var err error
//	maxAttempts := 10
//	databaseURL := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", cfg.DBUser, cfg.DBPassword, cfg.DBHost, cfg.DBPort, cfg.DBName)
//
//	for i := 0; i < maxAttempts; i++ {
//		log.Printf("connenting to %s", databaseURL)
//		db, err = sqlx.Open("mysql", databaseURL)
//		if err == nil {
//			err = db.Ping()
//			if err == nil {
//				log.Printf("Connected to database")
//				return db, nil
//			}
//		}
//		log.Printf("Failed to connect to database, retrying in 5 seconds...")
//		time.Sleep(5 * time.Second)
//	}
//	return nil, fmt.Errorf("failed to connect to database after 10 attempts")
//}
