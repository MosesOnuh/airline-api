package main

import (
	"log"
	"github.com/MosesOnuh/airline-api/config"
	"github.com/MosesOnuh/airline-api/Db/mongo"
	"github.com/MosesOnuh/airline-api/server"
	"fmt"
)

func main(){
	cfg := config.LoadConfig()


	_, _, err := Db.New(cfg.DBAddress, cfg.DBName, cfg.UserCollection, cfg.FlightCollection, cfg.TicketCollection)
	if err != nil {
		log.Fatalf("failed to open mongodb: %v", err)
	}
	
	go func() {
		addr := fmt.Sprintf(":%s", cfg.Port)
		err := server.Run(addr)
		if err != nil {
			log.Fatal(fmt.Sprintf("failed to start service: %v", err))
		}
	}()

}



