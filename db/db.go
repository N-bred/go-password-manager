package db

import (
	"database/sql"
	"log"
	"main/models"

	_ "github.com/ncruces/go-sqlite3/driver"
)

type DB struct {
	*sql.DB
}

func CreateDB(path string) (*DB, error) {

	db, err := sql.Open("sqlite3", path)

	if err != nil {
		return nil, err
	}

	Db := DB{
		db,
	}

	return &Db, nil
}

func (db *DB) PopulateDB(credentials *[]models.Credential) {

	_, err := db.Exec(createSchema())

	if err != nil {
		log.Fatal(err)
	}

	for _, credential := range *credentials {
		db.AddCredential(&credential)
	}
}

func createSchema() string {
	return `
	CREATE TABLE IF NOT EXISTS Credentials (
		id INTEGER PRIMARY KEY,
		domain TEXT NOT NULL,
		username TEXT NOT NULL,
		password TEXT NOT NULL
	);
	
	CREATE TABLE IF NOT EXISTS Credentials_History (
		id INTEGER PRIMARY KEY,
		credential_id INTEGER NOT NULL,
		value TEXT NOT NULL,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		FOREIGN KEY (credential_id) REFERENCES Credentials(id)
	);

	`
}

func (db *DB) AddCredential(credential *models.Credential) error {
	query := "INSERT INTO Credentials (domain, username, password) VALUES (?,?,?);"
	result, err := db.Exec(query, credential.Domain, credential.Username, credential.Password)

	if err != nil {
		return err
	}

	id, err := result.LastInsertId()

	if err != nil {
		return err
	}

	err = db.AddCredentialHistory(int(id), credential.Password)

	if err != nil {
		return err
	}

	return nil

}

func (db *DB) AddCredentialHistory(id int, password string) error {
	query := "INSERT INTO Credentials_History (credential_id, value) VALUES (?,?);"
	_, err := db.Exec(query, id, password)

	if err != nil {
		return err
	}

	return nil

}

func (db *DB) GetAllCredentials() ([]models.Credential, error) {
	query := "SELECT * FROM Credentials;"
	rows, err := db.Query(query)

	if err != nil {
		return nil, err
	}

	c := make([]models.Credential, 0)

	for rows.Next() {
		var cr models.Credential
		err := rows.Scan(&cr.Id, &cr.Domain, &cr.Username, &cr.Password)

		if err != nil {
			return nil, err
		}

		c = append(c, cr)
	}

	return c, nil
}

func (db *DB) GetCredentialById(id int) (*models.Credential, error) {
	query := "SELECT * FROM Credentials WHERE id = ?;"
	row := db.QueryRow(query, id)

	var c models.Credential

	err := row.Scan(&c.Id, &c.Domain, &c.Username, &c.Password)

	if err != nil {
		return nil, err
	}

	return &c, nil
}

func (db *DB) GetCredentialHistoryById(id int) (*[]models.CredentialHistory, error) {
	query := "SELECT * FROM Credentials_History WHERE credential_id = ? ORDER BY created_at DESC;"
	rows, err := db.Query(query, id)

	if err != nil {
		return nil, err
	}

	ch := make([]models.CredentialHistory, 0)

	for rows.Next() {
		var cr models.CredentialHistory
		err := rows.Scan(&cr.Id, &cr.Credential_id, &cr.Value, &cr.Created_at)

		if err != nil {
			return nil, err
		}

		ch = append(ch, cr)
	}

	return &ch, nil
}

func (db *DB) UpdateCredentialById(id int, credential *models.Credential) error {
	err := db.AddCredentialHistory(credential.Id, credential.Password)

	if err != nil {
		return err
	}

	query2 := "UPDATE Credentials SET domain = ?, username = ?, password = ? WHERE id = ?;"
	_, err = db.Exec(query2, credential.Domain, credential.Username, credential.Password, id)

	if err != nil {
		return err
	}

	return nil
}
