package server

import (
	"os"
	"github.com/MosesOnuh/airline-api/handlers"
	"github.com/gin-gonic/gin"
)

func Run( port string) error {
	router := gin.Default()
	router.POST("signupUser", handlers.SignupHandler)
	router.POST("loginUser", handlers.LoginHandler)
	router.POST("/createFlight", handlers.CreateFlightHandler)
	router.GET("/getFlight", handlers.CreateFlightHandler)
	router.PATCH("/updateflight", handlers.UpdateFlightHandler)
	//router.GET("/getAllFLight", handlers.GetAllFlightHandler)
	router.GET("/getSingleFLight", handlers.GetSingleFlightHandler)
	router.PATCH("/deleteFlight", handlers.DeleteFlightHandler)

	err := router.Run(":" + os.Getenv("PORT"))
		if err != nil {
				return err
			}
			return nil
}