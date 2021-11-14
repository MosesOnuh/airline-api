package server

import (
	"github.com/MosesOnuh/airline-api/handlers"
	"github.com/gin-gonic/gin"
)
 

func Run(Port string) error {
	h := &handlers.Handler{}
	router := gin.Default()
	router.POST("/signupUser", h.SignupHandler)
	router.POST("/loginUser", h.LoginHandler)
	router.POST("/createFlight", h.CreateFlightHandler)
	router.GET("/getFlight", h.CreateFlightHandler)
	router.PATCH("/updateflight", h.UpdateFlightHandler)
	router.GET("/getSingleFLight", h.GetSingleFlightHandler)
	router.PATCH("/deleteFlight", h.DeleteFlightHandler)
	
	err := router.Run(":" + Port)
	if err != nil {
		return err
	}
	return nil
}
