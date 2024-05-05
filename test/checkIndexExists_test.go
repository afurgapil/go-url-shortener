package helpers

import (
	"database/sql"
	"fmt"
	"log"
	"testing"

	"github.com/afurgapil/go-url-shortener/configs"
	"github.com/afurgapil/go-url-shortener/pkg/helpers"
	_ "github.com/go-sql-driver/mysql"
)
func createTestData(db *sql.DB)  {
	query := "INSERT INTO urls (long_url, short_url, usage_count, pass_key) VALUES (?, ?, ?, ?)"
	_, err := db.Exec(query, "https://www.example.com", "example", 0, "test_pass_key")
	if err != nil {
		log.Fatalf("Failed to insert value into test DB: %v", err)
	}
}

func clearTestData(db *sql.DB)  {
	query := "DELETE FROM urls"
	_, err := db.Exec(query)
	if err != nil {
		log.Fatalf("Failed to clear test data from test DB: %v", err)
	}
}

func TestCheckIndexExists(t *testing.T) {
	config, err := configs.Load()
	if err != nil {
		t.Fatalf("Error loading test config: %v", err)
	}

	db, err := sql.Open("mysql", config.DatabaseURL)
	if err != nil {
		t.Fatalf("Error connecting to test database: %v", err)
	}
	defer db.Close()

	createTestData(db)
	defer clearTestData(db)

	exists, err := helpers.CheckIndexExists("short_url", "example")
	fmt.Println(exists)
	if err != nil {
		t.Fatalf("Error checking index existence: %v", err)
	}

	if !exists {
		t.Error("Expected index to exist, got false")
	}
}
