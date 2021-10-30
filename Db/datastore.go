package Db

import "github.com/MosesOnuh/airline-api/models"

type Datastore interface {
	CreateUser(user *models.User) (*models.User, error)
	CheckUserExists(email string) bool 
	GetAllUsers() ([]models.User, error)
	GetUserByEmail(email string) (*models.User, error)
	CreateFlight (flight *models.Flight) (*models.Flight, error)
	GetAllFlight ()([]models.Flight, error)
	GetFlightByID (flightId string)(*models.Flight, error)
	UpdateFlight (
		flightId string,
		Country string,
		Departure_location string,
		Arrival_location string,
		Departure_time string,
		Arrival_time string,
		Price int,
		Available_seats int) error
	 DeleteFlight(flightId, AdminId string) error
}