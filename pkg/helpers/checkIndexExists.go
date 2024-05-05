package helpers

import (
	"database/sql"
	"log"

	"github.com/afurgapil/go-url-shortener/configs"
)


func CheckIndexExists(column, value string) (bool, error) {
	config, err := configs.Load()
	if err != nil {
		log.Fatalf("Error loading config: %v", err)
	}

	db, err := sql.Open("mysql", config.DatabaseURL)
	if err != nil {
		return false, err
	}
	defer db.Close()

	query := "SELECT COUNT(*) FROM urls WHERE `" + column + "` = ?"

	var count int
	err = db.QueryRow(query, value).Scan(&count)
	if err != nil {
		log.Println("Failed to execute query:", err)
		return false, err
	}

	return count > 0, nil
}

