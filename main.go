package main

import (
	"fmt"
	"log"
	"main/db"
	"main/routes"
	"net/http"

	"github.com/rs/cors"
)

func main() {
	db, err := db.GetDb()

	if err != nil {
		log.Fatal(err)

	}

	mux := http.NewServeMux()

	mux.HandleFunc("/credentials/", routes.GetAllCredentials(db))

	mux.HandleFunc("/credential/{id}", routes.GetCredentialById(db))

	mux.HandleFunc("/credential/create", routes.CreateCredential(db))

	mux.HandleFunc("/credential/update", routes.UpdateCredentialById(db))

	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "POST", "PUT"},
		Debug:          true,
	})

	handler := c.Handler(mux)

	fmt.Println("Listening on PORT: 3000")

	if err := http.ListenAndServe("127.0.0.1:3000", handler); err != nil {
		panic(err)
	}

}
