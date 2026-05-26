package main

import (
	"database/sql"
	"log"

	_ "github.com/ncruces/go-sqlite3/driver"
)

type DB struct {
	*sql.DB
}

func GetDb() (*DB, error) {
	db, err := sql.Open("sqlite3", "db.db")

	if err != nil {
		return nil, err

	}

	Db := DB{
		db,
	}

	_, err = Db.Exec(createTable())

	if err != nil {
		log.Fatal(err)
	}

	return &Db, nil
}

func createTable() string {
	return `CREATE TABLE IF NOT EXISTS Credentials (
		id INTEGER PRIMARY KEY,
		domain TEXT,
		username TEXT,
		password TEXT
	);`
}

func (db *DB) AddCredential(credential *Credential) {
	query := "INSERT INTO Credentials (domain, username, password) VALUES (?,?,?);"
	_, err := db.Exec(query, credential.Domain, credential.Username, credential.Password)

	if err != nil {
		log.Fatal(err)
	}

}

func (db *DB) GetAllCredentials() []Credential {
	query := "SELECT * FROM Credentials;"
	rows, err := db.Query(query)

	if err != nil {
		log.Fatal(err)
	}

	c := make([]Credential, 0)

	for rows.Next() {
		var cr Credential
		err := rows.Scan(&cr.Id, &cr.Domain, &cr.Username, &cr.Password)

		if err != nil {
			log.Fatal(err)
		}

		c = append(c, cr)
	}

	return c
}

func (db *DB) GetCredentialById(id int) *Credential {
	query := "SELECT * FROM Credentials WHERE id = ?;"
	row := db.QueryRow(query, id)

	var c Credential

	err := row.Scan(&c.Id, &c.Domain, &c.Username, &c.Password)

	if err != nil {
		log.Fatal(err)
	}

	return &c
}

func (db *DB) UpdateCredentialById(id int, credential *Credential) error {
	query := "UPDATE Credentials SET domain = ?, username = ?, password = ? WHERE id = ?;"
	_, err := db.Exec(query, credential.Domain, credential.Username, credential.Password, id)

	if err != nil {
		log.Fatal(err)
	}

	return nil
}
