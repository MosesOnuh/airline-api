package main

import (
	"os"
	"log"
	"github.com/MosesOnuh/airline-api/config"
	"github.com/MosesOnuh/airline-api/Db/mongo"
	"github.com/MosesOnuh/airline-api/server"
)

func main(){
	cfg := config.LoadConfig()

	mongoStore, _, err := Db.New(cfg.DBAddress, cfg.DBName, cfg.UserCollection, cfg.FlightCollection, cfg.TicketCollection)
	if err != nil {
		log.Fatalf("failed to open mongodb: %v", err)
	}
}

