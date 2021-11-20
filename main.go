package main

import (
	"fmt"
	"github.com/MosesOnuh/airline-api/auth"
	"github.com/MosesOnuh/airline-api/config"
	"github.com/MosesOnuh/airline-api/db/mongo"
	"github.com/MosesOnuh/airline-api/server"
	"log"
)

func main() {
	cfg := config.LoadConfig()

	datastore, _, err := db.New(cfg.DBAddress, cfg.DBName)
	if err != nil {
		log.Fatalf("failed to open mongodb: %v", err)
	}
	tokenHandler := auth.New(cfg.JWTSectret)

	addr := cfg.Port
	error := server.Run(addr, datastore, tokenHandler)
	if error != nil {
		log.Fatal(fmt.Sprintf("failed to start service: %v", err))
	}

}
