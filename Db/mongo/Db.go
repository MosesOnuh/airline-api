package Db
import (
	"context"
	"time"
	"github.com/MosesOnuh/airline-api/Db"
	"github.com/MosesOnuh/airline-api/models"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
    "go.mongodb.org/mongo-driver/mongo/readpref"
    "go.mongodb.org/mongo-driver/bson"
	 

)

//mongoStore, _, err := mongo.New(cfg.DBURL, cfg.DBName)

type mongoStore struct {
	client *mongo.Client
	dbName string
	userCollection string
	flightCollection string
	ticketCollection string
}

var _ Db.Datastore = &mongoStore{}


// defaultDbAddress        = "mongodb://localhost:27017"
// 	defaultDbName           = "airline"
// 	defaultUserCollection   = "users"
// 	defaultFlightCollection = "flights"
// 	defaultTicketCollection = "tickets"

//New returns an instance of mongo store
func New(dbAddress, dbName, userCollection, flightCollection, ticketCollection string) (Db.Datastore, *mongo.Client, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(dbAddress))

	if err != nil {
		return nil, nil, err
	}

	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		return nil, nil, err
		// log.Fatalf("Mongo db not available: %v\n", err)
	}

	return &mongoStore{
		client: client,
		dbName: dbName,
		userCollection: userCollection,
		flightCollection: flightCollection,
		ticketCollection: ticketCollection,
	}, client, nil
}


func (m *mongoStore) dbCol(collectionName string) *mongo.Collection {
	return m.client.Database(m.dbName).Collection(collectionName)
}


func (m mongoStore) CreateUser(user *models.User) (*models.User, error) {
	_, err := m.dbCol(m.userCollection).InsertOne(context.Background(), user)
	if err != nil {
		return nil, err
	}
	return user, nil
}


func (m mongoStore) CheckUserExists(email string) bool {

	query := bson.M{
		"email": email,
	}
	count, err := m.dbCol(m.userCollection).CountDocuments(context.Background(), query)
	if err != nil {
		return false
	}
	if count > 1 {
		return true
	}
	return false
}

func (m mongoStore) GetAllUsers() ([]models.User, error) {
	
	var users []models.User

	cursor, err := m.dbCol(m.userCollection).Find(context.Background(), bson.M{})
	if err != nil {
		return nil, err
	}
	err = cursor.All(context.Background(), &users)
	if err != nil {
		return nil, err
	}
	return users, nil
}

func(m mongoStore) GetUserByEmail(email string) (*models.User, error) {
	var user models.User
	query := bson.M{
		"email": email,
	}
	err := m.dbCol(m.userCollection).FindOne(context.Background(), query).Decode(&user)

	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (m mongoStore) CreateFlight (flight *models.Flight) (*models.Flight, error) {
	_, err := m.dbCol(m.flightCollection).InsertOne(context.Background(), flight)

	return flight, err
}
func (m mongoStore) GetAllFlight ()([]models.Flight, error){

	var flight []models.Flight

	cursor, err := m.dbCol(m.flightCollection).Find(context.Background(), bson.M{})
	if err != nil {
		return nil, err
	}
	err = cursor.All(context.Background(), &flight)
	if err != nil {
		return nil, err
	}
	return flight, nil
}

func (m mongoStore) GetFlightByID (flightId string)(*models.Flight, error){	
	var flight models.Flight
	 query := bson.M{
		"id": flightId,
	 }
	 err := m.dbCol(m.flightCollection).FindOne(context.Background(), query).Decode(&flight)
	 if err != nil {
		 return nil, err
	 } 
	 return &flight, nil
}

func (m mongoStore) UpdateFlight (
	flightId string,
	 Country string,
	 Departure_location string,
	 Arrival_location string,
	 Departure_time string,
	 Arrival_time string,
	 Price int,
	 Available_seats int) error {
		
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

	_, err := m.dbCol(m.flightCollection).UpdateOne(context.Background(),filterQuery, updateQuery)
	if err != nil {
		return err
	}
	return nil
}

func (m mongoStore) DeleteFlight(flightId, AdminId string) error {

	query := bson.M{
		"id": "flightId",
		"admin_id": "AdminId",
	}
	_, err := m.dbCol(m.flightCollection).DeleteOne(context.Background(), query)
	if err != nil {
		return err
	}
	return nil

}


// func CreateTicket (ticket *models.Ticket) (*models.Ticket, error) {
// 	_, err := DbClient.Database(os.Getenv("DB_NAME")).Collection(os.Getenv("TICKET_COLLECTION")).InsertOne(context.Background(), ticket)

// 	return ticket, err
// }
