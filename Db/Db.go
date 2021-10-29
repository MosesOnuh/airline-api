package Db
import (
	"log"
	"os"
	"context"
	"time"
	"github.com/MosesOnuh/airline-api/models"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
    "go.mongodb.org/mongo-driver/mongo/readpref"
    "go.mongodb.org/mongo-driver/bson"
	 "github.com/joho/godotenv"

)

var DbClient *mongo.Client
func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	mongoAddress :=os.Getenv("MONGO_ADDRESS")
	
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoAddress))

	if err != nil {
		log.Fatalf("Could not connect to the database %v\n", err)
	}
    DbClient = client
	err = DbClient.Ping(ctx, readpref.Primary())
	if err != nil {

		log.Fatalf("Mongo db not available: %v\n", err)
	}
}

func CreateUser(user *models.User) (*models.User, error){
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	dbName := os.Getenv("DB_NAME")
	userCollection := os.Getenv("USER_COLLECTION")


	_, err := DbClient.Database(dbName).Collection(userCollection).InsertOne(context.Background(), user)
	
	return user, err
}
func CheckUserExists(email string) bool {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	dbName := os.Getenv("DB_NAME")
	userCollection := os.Getenv("USER_COLLECTION")

	query := bson.M{
		"email": email,
	}
	count, err := DbClient.Database(dbName).Collection(userCollection).CountDocuments(context.Background(), query)
	if err != nil {
		return false
	}
	if count > 1 {
		return true
	}
	return false
}

func GetAllUsers() ([]models.User, error) {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	dbName := os.Getenv("DB_NAME")
	userCollection := os.Getenv("USER_COLLECTION")

	var users []models.User

	cursor, err := DbClient.Database(dbName).Collection(userCollection).Find(context.Background(), bson.M{})
	if err != nil {
		return nil, err
	}
	err = cursor.All(context.Background(), &users)
	if err != nil {
		return nil, err
	}
	return users, nil
}

func GetUserByEmail(email string) (*models.User, error) {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	dbName := os.Getenv("DB_NAME")
	userCollection := os.Getenv("USER_COLLECTION")

	var user models.User
	query := bson.M{
		"email": email,
	}
	err := DbClient.Database(dbName).Collection(userCollection).FindOne(context.Background(), query).Decode(&user)

	if err != nil {
		return nil, err
	}
	return &user, nil
}

func CreateFlight (flight *models.Flight) (*models.Flight, error) {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	dbName := os.Getenv("DB_NAME")
	userCollection := os.Getenv("USER_COLLECTION")

	_, err := DbClient.Database(dbName).Collection(userCollection).InsertOne(context.Background(), flight)

	return flight, err
}
func GetAllFlight ()([]models.Flight, error){
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	dbName := os.Getenv("DB_NAME")
	flightCollection := os.Getenv("FLIGHT_COLLECTION")

	var flight []models.Flight

	cursor, err := DbClient.Database(dbName).Collection(flightCollection).Find(context.Background(), bson.M{})
	if err != nil {
		return nil, err
	}
	err = cursor.All(context.Background(), &flight)
	if err != nil {
		return nil, err
	}
	return flight, nil
}

func GetFlightByID (flightId string)(*models.Flight, error){
 	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	dbName := os.Getenv("DB_NAME")
	flightCollection := os.Getenv("FLIGHT_COLLECTION")
	
	var flight models.Flight
	 query := bson.M{
		"id": flightId,
	 }
	 err := DbClient.Database(dbName).Collection(flightCollection).FindOne(context.Background(), query).Decode(&flight)
	 if err != nil {
		 return nil, err
	 } 
	 return &flight, nil
}

func UpdateFlight (
	flightId string,
	 Country string,
	 Departure_location string,
	 Arrival_location string,
	 Departure_time string,
	 Arrival_time string,
	 Price int,
	 Available_seats int) error {
		 err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	dbName := os.Getenv("DB_NAME")
	flightCollection := os.Getenv("FLIGHT_COLLECTION")

	filterQuery := bson.M{
	"id": flightId,
	}
	updateQuery := bson.M{
		"$set": bson.M{
			"country ": Country,
			"departure_location": Departure_location,
			"arrival_location": Arrival_location,
			"departure_time": Departure_time,
			"arrival_time": Arrival_time,
			"price": Price,
			"available_seats": Available_seats,
		},
	}

	_, err := DbClient.Database(dbName).Collection(flightCollection).UpdateOne(context.Background(),filterQuery, updateQuery)
	if err != nil {
		return err
	}
	return nil
}

func DeleteFlight(flightId, AdminId string) error {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	dbName := os.Getenv("DB_NAME")
	flightCollection := os.Getenv("FLIGHT_COLLECTION")

	query := bson.M{
		"id": "flightId",
		"admin_id": "AdminId",
	}
	_, err := DbClient.Database(dbName).Collection(flightCollection).DeleteOne(context.Background(), query)
	if err != nil {
		return err
	}
	return nil

}


// func CreateTicket (ticket *models.Ticket) (*models.Ticket, error) {
// 	_, err := DbClient.Database(os.Getenv("DB_NAME")).Collection(os.Getenv("TICKET_COLLECTION")).InsertOne(context.Background(), ticket)

// 	return ticket, err
// }
