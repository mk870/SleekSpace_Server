package main

import (
	"SleekSpace/controllers"
	"SleekSpace/db"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.New()
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowCredentials = true
	config.AddAllowHeaders("Authorization", "token", "User-Agent", "Accept")
	router.Use(cors.New(config))
	db.Connect()
	controllers.GetUsers(router)
	controllers.UpdateUser(router)
	controllers.GetUser(router)
	controllers.DeleteUser(router)
	controllers.Registration(router)
	controllers.Login(router)
	controllers.LoginOut(router)
	controllers.VerificationCodeValidation(router)
	controllers.ResendVerificationCode(router)
	router.Run()
}
