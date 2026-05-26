package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Credential struct {
	Domain   string
	Username string
	Password string
}

func (c *Credential) GetCredData() {
	fmt.Printf("Domain: %v; %v:%v\n", c.Domain, c.Username, c.Password)
}

func encodeCredentialsToJson(fileName string, credentials *[]Credential) {
	file, _ := os.Create(fileName)
	encoder := json.NewEncoder(file)
	encoder.SetEscapeHTML(false)
	encoder.Encode(credentials)
}
