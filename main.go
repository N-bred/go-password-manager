package main

import (
	"fmt"
	"log"
	"main/db"
	"main/routes"
	"net/http"
)

func main() {
	db, err := db.GetDb()

	if err != nil {
		log.Fatal(err)

	}

	http.HandleFunc("/credentials/", routes.GetAllCredentials(db))

	http.HandleFunc("/credential/{id}", routes.GetCredentialById(db))

	fmt.Println("Listening on PORT: 3000")

	if err := http.ListenAndServe("127.0.0.1:3000", nil); err != nil {
		panic(err)
	}

}
