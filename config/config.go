package config

import (
	"github.com/joho/godotenv"
	"os"
)

const (
	defaultPort             = "3000"
	defaultSecret           = "secret"
	defaultDbAddress        = "mongodb://localhost:27017"
	defaultDbName           = "airline"
)

// Configuration contains all the config that the appliction needs
type Configurations struct {
	Port             string `json:"port"`
	JWTSectret       string `json:"jwt_secret"`
	DBAddress        string `json:"db_address"`
	DBName           string `json:"db_name"`
}

func LoadConfig(filename ...string) *Configurations {
	e := ".env"
	if len(filename) > 0 {
		e = filename[0]
	}
	_ = godotenv.Load(e)
	configurations := &Configurations{}
	port, ok := os.LookupEnv("PORT")
	if !ok {
		port = defaultPort
	}
	configurations.Port = port

	jwtSecret, ok := os.LookupEnv("JWT_SECRET")
	if !ok {
		jwtSecret = defaultSecret
	}
	configurations.JWTSectret = jwtSecret

	dbAddress, ok := os.LookupEnv("MONGO_ADDRESS")
	if !ok {
		dbAddress = defaultDbAddress
	}
	configurations.DBAddress = dbAddress

	dbName, ok := os.LookupEnv("DB_NAME")
	if !ok {
		dbName = defaultDbName
	}
	configurations.DBName = dbName

	return configurations

}
