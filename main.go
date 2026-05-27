package main

import (
	"fmt"
	"log"
	dbi "main/db"
	"main/helpers"
	"main/routes"
	"net/http"
	"os"
	"slices"

	"github.com/rs/cors"
)

func main() {

	db, err := dbi.GetDb()

	if err != nil {
		log.Fatal(err)
	}

	if idx := slices.IndexFunc(os.Args, helpers.FindIdx("--populate")); idx != -1 {
		dbi.CreateAndPopulateDB()
	}

	mux := http.NewServeMux()

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
