package db

import (
	"database/sql"
	"log"

	"github.com/afurgapil/go-url-shortener/configs"
	"github.com/afurgapil/go-url-shortener/pkg/helpers"

	_ "github.com/go-sql-driver/mysql"
)

func PingDB(db *sql.DB) bool  {
    if err:=db.Ping(); err!=nil  {
        return false
    }
    return true
}


func NewDB(dataSourceName string) (*sql.DB, error) {
    db, err := sql.Open("mysql", dataSourceName)
    if err != nil {
        return nil, err
    }

    return db, nil
}

func CreateURL(longURL, shortURL string) (int64, error) {
    passKey := helpers.PassKeyGenerator(8)
    config, err := configs.Load()
	if err != nil {
		log.Fatalf("Error loading config: %v", err)
	}
	db, err := sql.Open("mysql", config.DatabaseURL)
	if err != nil {
		return 0, err
	}
	defer db.Close()

	query := "INSERT INTO urls (long_url, short_url,pass_key) VALUES (?, ?, ?)"

	result, err := db.Exec(query, longURL, shortURL, passKey)
	if err != nil {
		log.Println("Failed to insert URL into database:", err)
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		log.Println("Failed to get last insert ID:", err)
		return 0, err
	}

	return id, nil
}

func UseURL(column, value string) (string, error) {
	config, err := configs.Load()
	if err != nil {
		log.Fatalf("Error loading config: %v", err)
	}
		db, err := sql.Open("mysql", config.DatabaseURL)

	if err != nil {
		return "", err
	}
	defer db.Close()

	var passKey string
	query := "SELECT pass_key FROM urls WHERE " + column + " = ?"
	err = db.QueryRow(query, value).Scan(&passKey)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Println("No rows found")
		} else {
			log.Println("Error querying database:", err)
		}
		return "", err
	}

	return passKey, nil
}

func DeleteURL(column, value string) error {
	  config, err := configs.Load()
	if err != nil {
		log.Fatalf("Error loading config: %v", err)
	}
		db, err := sql.Open("mysql", config.DatabaseURL)

	if err != nil {
		return err
	}
	defer db.Close()

	query := "DELETE FROM urls WHERE " + column + " = ?"
	_, err = db.Exec(query, value)
	if err != nil {
		log.Println("Error deleting URL from database:", err)
		return err
	}

	return nil
}

func UpdateURL(column, value, newLongURL string) (string, error) {
	newPassKey := helpers.PassKeyGenerator(8)
	config, err := configs.Load()
	if err != nil {
		return "", err
	}

	db, err := sql.Open("mysql", config.DatabaseURL)
	if err != nil {
		return "", err
	}
	defer db.Close()

	query := "UPDATE urls SET long_url = ?, pass_key = ? WHERE " + column + " = ?"
	_, err = db.Exec(query, newLongURL, newPassKey, value)
	if err != nil {
		return "", err
	}

	return newLongURL, nil
}

func IncrementUsage(shortURL string) (bool, error) {
	config, err := configs.Load()
	if err != nil {
		return false, err
	}

	db, err := sql.Open("mysql", config.DatabaseURL)
	if err != nil {
		return false, err
	}
	defer db.Close()

	query := "UPDATE urls SET usage_count = usage_count + 1 WHERE short_url = ?"
	result, err := db.Exec(query, shortURL)
	if err != nil {
		return false, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return false, err
	}

	return rowsAffected > 0, nil
}


