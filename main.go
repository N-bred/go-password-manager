package main

import (
	"log"
)

func main() {

	credentials, err := TransformIntoCredentialStruct()

	if err != nil {
		log.Fatal("Found not Psx.txt file")
	}

	encodeCredentialsToJson("psx.json", &credentials)

}
