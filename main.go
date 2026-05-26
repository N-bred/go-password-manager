package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func transformData() *[]Credential {
	credentials, err := TransformIntoCredentialStruct()

	if err != nil {
		log.Fatal("Found not Psx.txt file")
	}

	encodeCredentialsToJson("psx.json", &credentials)

	return &credentials
}

func createAndPopulateDB() {
	credentials := transformData()

	db, err := GetDb()

	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	for _, credential := range *credentials {
		db.AddCredential(&credential)
	}
}

func main() {
	db, err := GetDb()

	if err != nil {
		log.Fatal(err)

	}

	http.HandleFunc("/credentials", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		credentials := db.GetAllCredentials()
		encoder := json.NewEncoder(w)
		encoder.SetEscapeHTML(false)
		encoder.Encode(credentials)
	})

	http.HandleFunc("/credential/{id}", func(w http.ResponseWriter, r *http.Request) {
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
	})

	fmt.Println("Listening on PORT: 3000")

	if err := http.ListenAndServe("127.0.0.1:3000", nil); err != nil {
		panic(err)
	}

}
