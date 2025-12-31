package main

import (
	"encoding/json"
	"log"
	"os"
)

type Settings struct {
	ExternalAddress string `json:"external_address"`
	InternalAddress string `json:"internal_address"`

	MaxRegisteredUsers uint `json:"max_registered_users"`
}

func main() {
	const SETTINGS_FILE_PATH = "./settings.authorization_microservice.json"

	settings := Settings{
		ExternalAddress: ":5202",
		InternalAddress: ":5302",

		MaxRegisteredUsers: 100,
	}

	content, err := os.ReadFile(SETTINGS_FILE_PATH)
	if err != nil {
		log.Println("Cannot read settings file:", err)

		content, err = json.Marshal(settings)
		if err != nil {
			log.Println("Cannot serialize default settings:", err)
		} else {
			err = os.WriteFile(SETTINGS_FILE_PATH, content, 0644)
			if err != nil {
				log.Println("Cannot write settings to file:", err)
			}
		}
	} else {
		err = json.Unmarshal(content, &settings)
		if err != nil {
			log.Println("Cannot deserialize settings:", err)
		}
	}
}
