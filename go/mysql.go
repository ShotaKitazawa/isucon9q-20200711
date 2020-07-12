package main

import (
	"log"
	"time"

	"github.com/jmoiron/sqlx"
)

func waitDB(db *sqlx.DB) {
	for {
		err := db.Ping()
		if err == nil {
			return
		}

		log.Printf("Failed to ping DB: %s", err)
		log.Println("Retrying...")
		time.Sleep(time.Second)
	}
}

func pollDB(db *sqlx.DB) {
	for {
		err := db.Ping()
		if err != nil {
			log.Printf("Failed to ping DB: %s", err)
		}

		time.Sleep(time.Second)
	}
}
