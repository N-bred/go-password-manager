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
	"path/filepath"
	"slices"
	"strings"

	"github.com/rs/cors"
)

func getAbsPath() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))

	if err != nil {
		log.Fatal(err)
	}

	return dir
}

//go:embed ui/dist/index.html
var html string

func main() {
	print(os.Args[1])
	files, err := os.ReadDir(os.Args[1])

	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		fmt.Printf("%v", file)
	}
}

func a() {

	db, err := dbi.CreateDB(filepath.Join(getAbsPath(), "db.db"))

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

		db.PopulateDB(&credentials)
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
