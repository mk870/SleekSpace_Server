package admin

import (
	"fmt"
	"net/http"
	"time"

	managerRepo "SleekSpace/repositories/manager"
	userRepo "SleekSpace/repositories/user"
	"SleekSpace/utilities"

	"github.com/gin-gonic/gin"
)

func GetLocationById(c *gin.Context) {
	id := c.Param("id")
	location := userRepo.GetLocationById(utilities.ConvertStringToInt(id))
	c.JSON(http.StatusOK, gin.H{"response": location})
}

func apiSim() <-chan string {
	result := make(chan string)
	go func() {
		defer close(result)
		time.Sleep(time.Second * 5)
		result <- "im done"
	}()
	return result
}

func GetInfo(c *gin.Context) {
	info := <-apiSim()
	fmt.Println("ran aftrer")
	c.JSON(http.StatusOK, gin.H{"response": info})

}

func GetAllUsersLocations(c *gin.Context) {
	codes := userRepo.GetAllUsersLocations()
	c.JSON(http.StatusOK, gin.H{
		"response": codes,
	})
}

func GetAllManagersProfilePictures(c *gin.Context) {
	pictures := managerRepo.GetAllManagersProfilePictures()
	c.JSON(http.StatusOK, gin.H{
		"response": pictures,
	})
}

func GetVerificationCodeById(c *gin.Context) {
	id := c.Param("id")
	code := userRepo.GetVerificationCodeById(id)
	c.JSON(http.StatusOK, gin.H{
		"response": code,
	})
}

func GetAllVerificationCodes(c *gin.Context) {
	codes := userRepo.AllVerificationCodes()
	c.JSON(http.StatusOK, gin.H{
		"response": codes,
	})
}

func DeleteVerificationCode(c *gin.Context) {
	id := c.Param("id")
	isDeleted := userRepo.DeleteVerficationCode(utilities.ConvertStringToInt(id))
	if isDeleted {
		c.JSON(http.StatusOK, gin.H{
			"response": "code deleted",
		})
	}
}

func GetUserContacts(c *gin.Context) {
	numbers := userRepo.GetAllUsersContactNumbers()
	c.JSON(http.StatusOK, gin.H{
		"response": numbers,
	})
}

func GetAllManagersContacts(c *gin.Context) {
	numbers := managerRepo.GetAllManagersContacts()
	c.JSON(http.StatusOK, gin.H{
		"response": numbers,
	})
}

func GetAllManagers(c *gin.Context) {
	managers := managerRepo.GetAllManagers()
	c.JSON(http.StatusOK, gin.H{
		"response": managers,
	})
}
