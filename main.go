package main

import (
	"encoding/json"
	"log"
	"net/smtp"
	"os"
)

type Settings struct {
	ExternalAddress string `json:"external_address"`
	InternalAddress string `json:"internal_address"`

	MaxRegisteredUsers uint `json:"max_registered_users"`

	SMTPUserName      string `json:"smtp_username"`
	SMTPPassword      string `json:"smtp_password"`
	SMTPHost          string `json:"smtp_host"`
	SMTPPort          string `json:"smtp_port"`
	SMTPSenderAddress string `json:"smtp_sender_address"`
}

func main() {
	const SETTINGS_FILE_PATH = "./settings.authorization_microservice.json"

	settings := Settings{
		ExternalAddress: ":5202",
		InternalAddress: ":5302",

		MaxRegisteredUsers: 100,

		SMTPUserName:      "Admin",
		SMTPPassword:      "Admin",
		SMTPHost:          "localhost",
		SMTPPort:          "587",
		SMTPSenderAddress: "xenon_authorization_nicroservice@example.com",
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
		} else {
			log.Println("Settings successfully deserialized.")
		}
	}

	auth := smtp.PlainAuth("", settings.SMTPUserName, settings.SMTPPassword, settings.SMTPHost)

	to := "discordeg@mail.ru"
	msg := []byte("To: " + to + "\r\n" +
		"Subject: TEST SUBJECT\r\n" +
		"\r\n" +
		"This is the email body.\r\n")
	err = smtp.SendMail(settings.SMTPHost+":"+settings.SMTPPort, auth, settings.SMTPSenderAddress, []string{to}, msg)
	if err != nil {
		log.Fatal("Cannot send mail:", err)
	}

	log.Println("Message sent successfully!")
}
