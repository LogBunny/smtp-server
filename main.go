package main

import (
	"log"
	"smtpserver/config"
	"smtpserver/db"
	"smtpserver/migrations"
	"smtpserver/src/backend"
	"smtpserver/utils"
	"time"

	"github.com/emersion/go-smtp"
)

func deleteOldMails() {
	ticker := time.NewTicker(12 * time.Hour)
	defer ticker.Stop()

	for {
		<-ticker.C // Wait for the next tick
		err := db.EmailSvc.DeleteEmails()
		if err != nil {
			log.Printf("Failed to delete old mails: %v", err)
		} else {
			log.Println("Old mails deleted successfully.")
		}
	}
}

func main() {
	utils.ImportEnv()
	config.LoadCfg()
	db.InitServices()
	if config.MIGRATE {
		migrations.Migrate()
	}
	go deleteOldMails()
	s := smtp.NewServer(&backend.SmtpBackend{})
	s.Addr = ":2525"
	s.Domain = "indiedev.blog"
	s.WriteTimeout = 10 * time.Second
	s.ReadTimeout = 10 * time.Second
	s.MaxMessageBytes = 1024 * 1024
	s.MaxRecipients = 50
	s.AllowInsecureAuth = true

	log.Println("Starting server at: ", s.Addr)
	if err := s.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
