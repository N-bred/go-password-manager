package models

import "fmt"

type CredentialHistory struct {
	Id            int    `json:"id"`
	Credential_id string `json:"credential_id"`
	Value         string `json:"value"`
	Created_at    string `json:"created_at"`
}

func (c *CredentialHistory) Print() {
	fmt.Printf("Id: %v; Cred Id: %v; Val: %v; Date: %v\n", c.Id, c.Credential_id, c.Value, c.Created_at)
}
