package helpers

import (
	"encoding/json"
	"log"
	"main/models"
	"os"
	"strings"
)

func TransformIntoCredentialStruct(filename string) ([]models.Credential, error) {
	data, err := os.ReadFile(filename)

	if err != nil {
		return nil, err
	}

	contents := strings.Split(strings.TrimSpace(string(data)), "\n")

	counter := 0

	credentials := make([]models.Credential, 0)

	for i, line := range contents {

		if line == "\r" {

			domain := strings.ReplaceAll(contents[i-3], "\r", "")
			username := strings.ReplaceAll(contents[i-2], "\r", "")
			password := strings.ReplaceAll(contents[i-1], "\r", "")

			if counter == 2 {
				domain = strings.ReplaceAll(contents[i-2], "\r", "")
				username = "thehardyboiz1@gmail.com"
			}

			cred := models.Credential{
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

func TransformDataIntoJson(filename string) *[]models.Credential {
	credentials, err := TransformIntoCredentialStruct(filename)

	if err != nil {
		log.Fatalf("Found no %v file", filename)
	}

	EncodeCredentialsToJson("p.json", &credentials)

	return &credentials
}

func TransformJsonIntoData(filename string) *[]models.Credential {
	content, err := os.ReadFile(filename)

	if err != nil {
		log.Fatalf("Couldn't read file: %v", filename)
	}

	var jsonContent []models.Credential
	err = json.Unmarshal(content, &jsonContent)

	if err != nil {
		log.Fatal("Unable to decode json file")
	}

	return &jsonContent
}
