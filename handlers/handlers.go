package handlers

import (
	"fmt"
	"strings"
	"time"

	"github.com/MosesOnuh/airline-api/db" 
	"github.com/MosesOnuh/airline-api/auth"
	"github.com/MosesOnuh/airline-api/models"
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
	
)

type Handler struct {
	Store db.Datastore
	JwtHandler auth.TokenHandler
}




func (h Handler) SignupHandler(c *gin.Context) {
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
	userCheck := h.Store.CheckUserExists(signupReq.Email)
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

	_, err = h.Store.CreateUser(&user)
	if err != nil {
		fmt.Println("error saving user", err)
		c.JSON(500, gin.H{
			"error": "Could not create user",
		})
		return
	}
	jwtTokenString, err := h.JwtHandler.CreateToken(user.ID)
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

func (h Handler) LoginHandler(c *gin.Context) {
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
	user, err := h.Store.GetUserByEmail(loginReq.Email)
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
	jwtTokenString, err := h.JwtHandler.CreateToken(user.ID)
	if err != nil {
		c.JSON(500, gin.H{
			"error": "invalid token",
		})
	}
	
	c.JSON(200, gin.H{
		"message": "Login successful",
		"token":   jwtTokenString,
		"data":    user,
	})
}

func (h Handler) GetAllUserHandler(c *gin.Context) {
	users, err := h.Store.GetAllUsers()
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


func (h Handler) CreateFlightHandler(c *gin.Context) {
	authorization := c.Request.Header.Get("Authorization")
	if authorization == "" {
		c.JSON(401, gin.H{
			"error": "token not supplied",
		})
		return
	}

	payload := ""
	splitTokenArray := strings.Split(authorization, " ")
	if len(splitTokenArray) > 1 {
		payload = splitTokenArray[1]
	}
	claims, err := h.JwtHandler.ValidToken(payload)
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
		Owner:              claims.UserId,
		Departure_location: flightDetails.Departure_location,
		Arrival_location:   flightDetails.Arrival_location,
		Departure_time:     flightDetails.Departure_time,
		Arrival_time:       flightDetails.Arrival_time,
		Price:              flightDetails.Price,	
	}

	_, err = h.Store.CreateFlight(&flight)
	if err != nil {
		fmt.Println("error saving task", err)
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

func (h Handler) GetAllFlightsHandler(c *gin.Context){
	authorization := c.Request.Header.Get("Authorization")
	if authorization == "" {
		c.JSON(401, gin.H{
			"error": "auth token not supplied",
		})
		return
	}

	payload := ""
	splitTokenArray := strings.Split(authorization, " ")
	if len(splitTokenArray) > 1 {
		payload = splitTokenArray[1]
	}
	claims, err := h.JwtHandler.ValidToken(payload)
	if err != nil {
		c.JSON(401, gin.H{
			"error": "invalid jwt token",
		})
		return
	}

	flights, err := h.Store.GetAllFlight(claims.UserId)
	if err != nil {
		c.JSON(500, gin.H{
			"error": "Could not process request, could get tasks",
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "success",
		"data":    flights,
    })
}

func (h Handler) GetSingleFlightHandler(c *gin.Context) {
	taskId := c.Param("id")
	task, err := h.Store.GetFlightByID(taskId)
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

func (h Handler) UpdateFlightHandler(c *gin.Context) {
	authorization := c.Request.Header.Get("Authorization")
	if authorization == "" {
		c.JSON(401, gin.H{
			"error": "auth token not supplied",
		})
		return
	}

	payload := ""
	splitTokenArray := strings.Split(authorization, " ")
	if len(splitTokenArray) > 1 {
		payload = splitTokenArray[1]
	}
	claims, error := h.JwtHandler.ValidToken(payload)
	if error != nil {
		c.JSON(401, gin.H{
			"error": "invalid jwt token",
		})
		return
	}
	owner := claims.UserId
	
	flightId := c.Param("id")

	var flight models.Flight

	err := c.ShouldBindJSON(&flight)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "invalid request data",
		})
		return
	}
	err = h.Store.UpdateFlight(flightId, owner, flight.Country, flight.Departure_location, flight.Arrival_location, flight.Departure_time, flight.Arrival_time, flight.Price)
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

func (h Handler) DeleteFlightHandler(c *gin.Context) {
	authorization := c.Request.Header.Get("Authorization")
	if authorization == "" {
		c.JSON(401, gin.H{
			"error": "auth token not supplied",
		})
		return
	}

	payload := ""
	splitTokenArray := strings.Split(authorization, " ")
	if len(splitTokenArray) > 1 {
		payload = splitTokenArray[1]
	}
	claims, err := h.JwtHandler.ValidToken(payload)
	if err != nil {
		c.JSON(401, gin.H{
			"error": "invalid jwt token",
		})
		return
	}

	flightId := c.Param("id")

	err = h.Store.DeleteFlight(flightId, claims.UserId)
	if err != nil {
		c.JSON(500, gin.H{
			"error": "flight could not be deleted",
		})
		}
	c.JSON(200, gin.H{
			"message": "Task deleted",
		})
	

}

