package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-openapi/runtime/middleware"
	"github.com/gorilla/mux"

	"github.com/afurgapil/go-url-shortener/configs"
	"github.com/afurgapil/go-url-shortener/internal/db"
	"github.com/afurgapil/go-url-shortener/internal/handlers"
)

func main() {
	config, err := configs.Load()
	if err != nil {
		log.Fatalf("Error loading config: %v", err)
	}

	fmt.Println("Veritabanı URL'si:", config.DatabaseURL)
	fmt.Println("Port:", config.Port)

	database, err := db.NewDB(config.DatabaseURL)
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}
	defer database.Close()

	if db.PingDB(database) {
		fmt.Println("Veritabanı ping isteği başarılı. Veritabanı erişimi başarılı bir şekilde yapıldı.")
	} else {
		log.Fatal("Veritabanı ping isteği başarısız. Veritabanı erişimi sağlanamadı.")
	}
	fmt.Println("Server is running on localhost:" + config.Port)

	r := mux.NewRouter()
	r.HandleFunc("/short", handlers.URLShorter)
	r.HandleFunc("/delete/{short_url}/{pass_key}", handlers.URLDeleter)
	r.HandleFunc("/update", handlers.URLUpdate)
	r.HandleFunc("/log/{id}", handlers.URLLogger)

	r.PathPrefix("/swagger.yaml").Handler(http.FileServer(http.Dir("./")))
	opts := middleware.SwaggerUIOpts{SpecURL: "swagger.yaml"}
	sh := middleware.SwaggerUI(opts, nil)
	r.Handle("/docs", sh)

	opts1 := middleware.RedocOpts{SpecURL: "swagger.yaml", Path: "doc"}
	sh1 := middleware.Redoc(opts1, nil)
	r.Handle("/doc", sh1)

	log.Fatal(http.ListenAndServe(":"+config.Port, r))
}
