package server

import (
	"github.com/MosesOnuh/airline-api/auth"
	"github.com/MosesOnuh/airline-api/db"
	"github.com/MosesOnuh/airline-api/handlers"
	"github.com/gin-gonic/gin"
)

func Run(Port string, datastore db.Datastore, secret auth.TokenHandler) error {
	h := &handlers.Handler{
		Store:      datastore,
		JwtHandler: secret,
	}
	router := gin.Default()
	router.POST("/signupUser", h.SignupHandler)
	router.POST("/loginUser", h.LoginHandler)
	router.GET("/getUsers", h.GetAllUserHandler)
	router.POST("/createFlight", h.CreateFlightHandler)
	router.GET("/getSingleFLight/:id", h.GetSingleFlightHandler)
	router.GET("/getAllFLights", h.GetAllFlightsHandler)
	router.PATCH("/updateflight/:id", h.UpdateFlightHandler)
	router.DELETE("/deleteFlight/:id", h.DeleteFlightHandler)

	err := router.Run(":" + Port)
	if err != nil {
		return err
	}
	return nil
}
