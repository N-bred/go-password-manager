package routes

import (
	"encoding/json"
	"log"
	"main/db"
	"net/http"
	"strconv"
)

func GetCredentialById(db *db.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		id, err := strconv.Atoi(r.PathValue("id"))

		if err != nil {

			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(`{"error": "Received non-integer in URL"}`))
			return
		}

		credential := db.GetCredentialById(int(id))

		if credential == nil {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte(`{"error": "Credential not found"}`))
			return
		}

		w.WriteHeader(http.StatusOK)
		encoder := json.NewEncoder(w)
		encoder.SetEscapeHTML(false)

		if err := encoder.Encode(credential); err != nil {
			log.Println("JSON encoding error:", err)
		}
	}
}
