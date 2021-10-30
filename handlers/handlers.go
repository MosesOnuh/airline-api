package handlers

import (
	"fmt"
	"strings"
	"time"

	"github.com/MosesOnuh/airline-api/Db"
	"github.com/MosesOnuh/airline-api/auth"
	"github.com/MosesOnuh/airline-api/models"
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
)

type handler struct {
	store Db.Datastore
}

func (h *handler) SignupHandler(c *gin.Context) {
	type SignupRequest struct {
		Name     string `json:"name"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	var signupReq SignupRequest

	err := c.ShouldBindJSON(&signupReq)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "Invalid request data",
		})
		return
	}
	userCheck := h.store.CheckUserExists(signupReq.Email)
	if userCheck {
		c.JSON(500, gin.H{
			"error": "User already exists, use another email",
		})
		return
	}

	bytes, err := bcrypt.GenerateFromPassword([]byte(signupReq.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(500, gin.H{
			"error": "Could not process request",
		})
		return
	}
	hashPassword := string(bytes)

	userId := uuid.NewV4().String()
	user := models.User{
		ID:       userId,
		Name:     signupReq.Name,
		Email:    signupReq.Email,
		Password: hashPassword,
		TS:       time.Now(),
	}

	_, err = h.store.CreateUser(&user)
	if err != nil {
		fmt.Println("error saving user", err)
		c.JSON(500, gin.H{
			"error": "Could not create user",
		})
		return
	}
	jwtTokenString, err := auth.CreateToken(user.ID)
	if err != nil {
		c.JSON(500, gin.H{
			"error": "invalid token",
		})
	}

	c.JSON(200, gin.H{
		"message": "sign up succesful",
		"token":   jwtTokenString,
		"data":    user,
	})
}

func (h handler) LoginHandler(c *gin.Context) {
	type loginDetails struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	var loginReq loginDetails

	err := c.ShouldBindJSON(&loginReq)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "Invalid Request Data",
		})
		return
	}
	user, err := h.store.GetUserByEmail(loginReq.Email)
	if err != nil {
		fmt.Printf("error getting user from dn: %v\n", err)
		c.JSON(500, gin.H{
			"error": "Could not process request, could get user",
		})
		return
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginReq.Password))
	if err != nil {
		fmt.Printf("error validating password: %v\n", err)
		c.JSON(500, gin.H{
			"error": "invalid login details",
		})
		return
	}
	c.JSON(200, gin.H{
		"message": "sign up successful",
		"token":   "jwtToken",
		"data":    user,
	})
}

func (h handler) GetAllUserHandler(c *gin.Context) {
	users, err := h.store.GetAllUsers()
	if err != nil {
		c.JSON(500, gin.H{
			"error": "User not found",
		})
		return
	}
	c.JSON(200, gin.H{
		"message": "success",
		"data":    users,
	})
}

// getsingleUser
// Post flight
// getallFlight
// getSingleflight

func (h handler) CreateFlightHandler(c *gin.Context) {
	authorization := c.Request.Header.Get("Authorization")
	if authorization == "" {
		c.JSON(401, gin.H{
			"error": "token not supplied",
		})
		return
	}

	payload := ""
	splitTokenArray := strings.Split(authorization, "")
	if len(splitTokenArray) > 1 {
		payload = splitTokenArray[1]
	}
	claims, err := auth.ValidToken(payload)
	if err != nil {
		c.JSON(401, gin.H{
			"error": "invalid jwt token",
		})
		return
	}
	var flightDetails models.Flight
	err = c.ShouldBindJSON(&flightDetails)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "Invalid Request Data",
		})
		return
	}
	flightId := uuid.NewV4().String()

	flight := models.Flight{
		ID:                 flightId,
		Country:            flightDetails.Country,
		Admin_Id:           claims.Id,
		Departure_location: flightDetails.Departure_location,
		Arrival_location:   flightDetails.Arrival_location,
		Departure_time:     flightDetails.Departure_time,
		Arrival_time:       flightDetails.Arrival_time,
		Price:              flightDetails.Price,
		Available_seats:    flightDetails.Available_seats,
	}

	_, err = h.store.CreateFlight(&flight)
	if err != nil {
		c.JSON(500, gin.H{
			"error": "Could not process request, unsaved flight data",
		})
		return
	}
	c.JSON(200, gin.H{
		"message": "successfully created flight",
		"data":    flight,
	})
}

func (h handler) GetSingleFlightHandler(c *gin.Context) {
	taskId := c.Param("id")
	task, err := h.store.GetFlightByID(taskId)
	if err != nil {
		c.JSON(404, gin.H{
			"error": "invalid task id" + taskId,
		})
		return
	}
	c.JSON(200, gin.H{
		"message": "success",
		"data":    task,
	})
}

func (h handler) UpdateFlightHandler(c *gin.Context) {
	flightId := c.Param("id")

	var flight models.Flight
	err := c.ShouldBindJSON(&flight)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "invalid request data",
		})
		return
	}
	err = h.store.UpdateFlight(flightId, flight.Country, flight.Departure_location, flight.Arrival_location, flight.Departure_time, flight.Arrival_time, flight.Price, flight.Available_seats)
	if err != nil {
		c.JSON(500, gin.H{
			"error": "Flight could not be updated",
		})
		return
	}
	c.JSON(200, gin.H{
		"message": "Task updated",
	})
}

func (h handler) DeleteFlightHandler(c *gin.Context) {
	authorization := c.Request.Header.Get("Authorization")
	if authorization == "" {
		c.JSON(401, gin.H{
			"error": "auth token not supplied",
		})
		return
	}

	payload := ""
	splitTokenArray := strings.Split(authorization, "")
	if len(splitTokenArray) > 1 {
		payload = splitTokenArray[1]
	}
	claims, err := auth.ValidToken(payload)
	if err != nil {
		c.JSON(401, gin.H{
			"error": "invalid jwt token",
		})
		return
	}

	flightId := c.Param("id")

	err = h.store.DeleteFlight(flightId, claims.UserId)
	if err != nil {
		c.JSON(500, gin.H{
			"error": "flight could not be deleted",
		})
		c.JSON(200, gin.H{
			"message": "Task deleted",
		})
	}

}
