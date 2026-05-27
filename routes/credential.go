package routes

import (
	"encoding/json"
	"log"
	"main/db"
	"main/models"
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

		credential, err := db.GetCredentialById(id)

		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(`{"error": "An error ocurred"}`))
			return
		}

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

func UpdateCredentialById(db *db.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		if r.Method != "PUT" {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(`{"error": "This endpoint only accepts PUT requests"}`))
			return
		}

		var c models.Credential

		err := json.NewDecoder(r.Body).Decode(&c)

		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(`{"error":"Invalid JSON Payload"}`))
			return
		}

		credential, err := db.GetCredentialById(c.Id)

		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(`{"error": "An error ocurred"}`))
			return
		}

		if credential == nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(`{"error":"Credential not found"}`))
			return
		}

		credential.Domain = c.Domain
		credential.Username = c.Username
		credential.Password = c.Password

		err = db.UpdateCredentialById(c.Id, credential)

		if err != nil {
			log.Fatal(err)
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(`{"error":"Something went wrong"}`))
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"response":"ok"}`))

	}
}

func CreateCredential(db *db.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		if r.Method != "POST" {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(`{"error": "This endpoint only accepts POST requests"}`))
			return
		}

		var c models.Credential

		err := json.NewDecoder(r.Body).Decode(&c)

		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(`{"error":"Invalid JSON Payload"}`))
			return
		}

		err = db.AddCredential(&c)

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(`{"error":"An error ocurred while trying to create the credential."}`))
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"response":"ok"}`))
	}
}
