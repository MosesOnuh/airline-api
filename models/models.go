package models

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

type User struct {
	ID               string    `json:"id" bson:"id"`
	Name             string    `json:"name" bson:"name"`
	Email            string    `json:"email" bson:"email"`
	Password         string    `json:"password" bson:"password"`
	Phone_No         string    `json:"phone_number" bson:"phone_no"`
	Gender           string    `json:"gender" bson:"gender"`
	Covid_Vac_Status string    `json:"covid_vac_status" bson:"covid_vac_status"`
	TS               time.Time `json:"timestamp" bson:"timestamp"`
	Purchased_Ticket int       `json:"purchased_ticket" bson:"purchased_ticket"`
}

type Flight struct {
	ID                 string `json:"id" bson:"id"`
	Admin_Id                 string `json:"admin_id" bson:"admin_id"`
	Country            string `json:"country" bson:"country"`
	Departure_location string `json:"departure_location" bson:"departure_location"`
	Arrival_location   string `json:"arrival_location" bson:"arrival_location"`
	Departure_time     string `json:"departure_time" bson:"departure_time"`
	Arrival_time       string `json:"arrival_time" bson:"arrival_time"`
	Price              int    `json:"price" bson:"price"`
	Available_Seats    int    `json:"available_seats" bson:"available_seats"`
}

type Ticket struct {
	ID             string `json:"id" bson:"id"`
	User_name      string `json:"user_name" bson:"user_name"`
	Departure_time string `json:"departure_time" bson:"departure_time"`
	Arrival_time   string `json:"arrival_time" bson:"arrival_time"`
	Price          int    `json:"price" bson:"price"`
}

type Claims struct {
	UserId string `json:"user_id"`
	jwt.StandardClaims
}
