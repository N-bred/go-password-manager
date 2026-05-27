package helpers

import (
	"encoding/json"
	"main/models"
	"os"
)

func EncodeCredentialsToJson(fileName string, credentials *[]models.Credential) {
	file, _ := os.Create(fileName)
	encoder := json.NewEncoder(file)
	encoder.SetEscapeHTML(false)
	encoder.Encode(credentials)
}
