package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/afurgapil/go-url-shortener/internal/db"
	"github.com/afurgapil/go-url-shortener/pkg/helpers"
)

func URLShorter(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Only POST requests are allowed", http.StatusMethodNotAllowed)
		return
	}

	var requestData struct {
		LongURL string `json:"long_url"`
	}
	if err := json.NewDecoder(r.Body).Decode(&requestData); err != nil {
		http.Error(w, "Invalid JSON body", http.StatusBadRequest)
		return
	}

	shortURL := helpers.URLShorter(requestData.LongURL)

	id, err := db.CreateURL(requestData.LongURL, shortURL)
	if err != nil {
		http.Error(w, "Failed to create shortened URL", http.StatusInternalServerError)
		return
	}

	responseData := struct {
		ID       int    `json:"id"`
		ShortURL string `json:"short_url"`
	}{ID: int(id), ShortURL: shortURL}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(responseData)
}


func URLDeleter(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		http.Error(w, "Only DELETE requests are allowed", http.StatusMethodNotAllowed)
		return
	}
	
	params := mux.Vars(r)
	shortURL := params["short_url"]
	passKey := params["pass_key"] 
	completedURL, err := helpers.UrlCompleter(shortURL)
	if err != nil {
		log.Println("Error completing URL:", err)
		http.Error(w, "Failed to complete URL", http.StatusInternalServerError)
		return
	}

	if completedURL == ""  {
		http.Error(w, "Missing short_url parameter", http.StatusBadRequest)
		return
	}

	if passKey == "" {
		http.Error(w, "Missing pass_key parameter", http.StatusBadRequest)
		return
	}

	expectedPassKey, err := db.UseURL("short_url", completedURL)
	if err != nil {
		http.Error(w, "Failed to use URL", http.StatusInternalServerError)
		return
	}

	if passKey == expectedPassKey {
	err := db.DeleteURL("short_url", completedURL)
	if err != nil {
		http.Error(w, "Failed to delete URL", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"message": completedURL + " deletes successfully.",
	})

	} else {
		http.Error(w, "Invalid pass key", http.StatusUnauthorized)
		return
	}

}

func URLUpdate(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPatch {
		http.Error(w, "Only PATCH requests are allowed", http.StatusMethodNotAllowed)
		return
	}

	var requestData struct {
		ShortURL   string `json:"shortURL"`
		NewLongURL string `json:"newLongURL"`
		PassKey    string `json:"passKey"`
	}
	if err := json.NewDecoder(r.Body).Decode(&requestData); err != nil {
		http.Error(w, "Invalid JSON body", http.StatusBadRequest)
		return
	}

	shortURL := requestData.ShortURL
	completedURL, err := helpers.UrlCompleter(shortURL)
	if err != nil {
		log.Println("Error completing URL:", err)
		http.Error(w, "Failed to complete URL", http.StatusInternalServerError)
		return
	}

	expectedPassKey, err := db.UseURL("short_url", completedURL)
	if err != nil {
		http.Error(w, "Failed to use URL", http.StatusInternalServerError)
		return
	}

	if expectedPassKey == "" {
    http.Error(w, "URL not found", http.StatusNotFound)
    return
	}


	if requestData.PassKey == expectedPassKey {
		updatedLongURL, err := db.UpdateURL("short_url", completedURL, requestData.NewLongURL)
		if err != nil {
			http.Error(w, "Failed to update URL", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		responseData := struct {
			Message     string `json:"message"`
			UpdatedURL  string `json:"updatedURL"`
		}{
			Message:     "URL updated successfully",
			UpdatedURL:  updatedLongURL,
		}
		json.NewEncoder(w).Encode(responseData)
	} else {
		http.Error(w, "Invalid pass key", http.StatusUnauthorized)
	}
}

func URLLogger(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Only GET requests are allowed", http.StatusMethodNotAllowed)
		return
	}

	params := mux.Vars(r)
	shortedURL := params["id"]

	completedURL, err := helpers.UrlCompleter(shortedURL)
	if err != nil {
		log.Println("Error completing URL:", err)
		http.Error(w, "Failed to complete URL", http.StatusInternalServerError)
		return
	}

	isExists, err := helpers.CheckIndexExists("short_url", completedURL)
	if err != nil {
		log.Println("Error checking URL existence:", err)
		http.Error(w, "Failed to check URL existence", http.StatusInternalServerError)
		return
	}
	if !isExists {
		http.Error(w, "URL not found", http.StatusNotFound)
		return
	}

	success, err := db.IncrementUsage(completedURL)
	if err != nil {
		log.Println("Error incrementing usage:", err)
		http.Error(w, "Failed to increment usage", http.StatusInternalServerError)
		return
	}

	if !success {
		log.Println("Failed to increment usage")
		http.Error(w, "Failed to increment usage", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	responseData := struct {
		Message     string `json:"message"`
		CompletedURL string `json:"completed_url"`
	}{
		Message:     "Usage incremented",
		CompletedURL: completedURL,
	}
	json.NewEncoder(w).Encode(responseData)
}

