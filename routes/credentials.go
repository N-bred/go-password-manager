package routes

import (
	"encoding/json"
	"log"
	"main/db"
	"net/http"
)

func GetAllCredentials(db *db.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		credentials, err := db.GetAllCredentials()

		if err != nil {
			log.Fatal(err)
		}

		encoder := json.NewEncoder(w)
		encoder.SetEscapeHTML(false)
		encoder.Encode(credentials)
	}
}
