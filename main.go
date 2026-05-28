package main

import (
	_ "embed"
	"fmt"
	"log"
	dbi "main/db"
	"main/helpers"
	"main/models"
	"main/routes"
	"net/http"
	"os"
	"slices"
	"strings"

	"github.com/rs/cors"
)

//go:embed ui/dist/index.html
var html string

func main() {

	db, err := dbi.GetDb()

	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	if idx := slices.IndexFunc(os.Args, helpers.FindIdx("--source")); idx != -1 {

		filename := os.Args[idx+1]

		var credentials []models.Credential

		if strings.Contains(filename, ".txt") {
			credentials = *helpers.TransformDataIntoJson(filename)
		} else if strings.Contains(filename, ".json") {
			credentials = *helpers.TransformJsonIntoData(filename)
		} else {
			log.Fatal("Passed an unsupported file source")
		}

		dbi.CreateAndPopulateDB(&credentials)
	}

	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		w.Write([]byte(html))
	})

	mux.HandleFunc("/credentials/", routes.GetAllCredentials(db))

	mux.HandleFunc("/credential/{id}", routes.GetCredentialById(db))

	mux.HandleFunc("/credential/create", routes.CreateCredential(db))

	mux.HandleFunc("/credential/update", routes.UpdateCredentialById(db))

	mux.HandleFunc("/credential-history/{id}", routes.GetCredentialHistoryById(db))

	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "POST", "PUT"},
		Debug:          false,
	})

	handler := c.Handler(mux)

	fmt.Println("Listening on PORT: 3000")

	if err := http.ListenAndServe("127.0.0.1:3000", handler); err != nil {
		panic(err)
	}

}
