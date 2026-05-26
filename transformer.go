package main

import (
	"os"
	"strings"
)

func TransformIntoCredentialStruct() ([]Credential, error) {
	data, err := os.ReadFile("Psx.txt")

	if err != nil {
		return nil, err
	}

	contents := strings.Split(strings.TrimSpace(string(data)), "\n")

	counter := 0

	credentials := make([]Credential, 0)

	for i, line := range contents {

		if line == "\r" {

			domain := strings.ReplaceAll(contents[i-3], "\r", "")
			username := strings.ReplaceAll(contents[i-2], "\r", "")
			password := strings.ReplaceAll(contents[i-1], "\r", "")

			if counter == 2 {
				domain = strings.ReplaceAll(contents[i-2], "\r", "")
				username = "thehardyboiz1@gmail.com"
			}

			cred := Credential{
				Domain:   domain,
				Username: username,
				Password: password,
			}

			credentials = append(credentials, cred)
			counter = -1
		}

		counter += 1

	}

	return credentials, nil
}
