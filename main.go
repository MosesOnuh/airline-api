package main

import (
	"os"
	"log"
	"github.com/MosesOnuh/airline-api/server"
)

func main(){
	err := server.Run(os.Getenv("PORT"))
	if err != nil {
		log.Fatal("Could not start server")
	}
}