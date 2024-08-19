package property

import (
	"net/http"

	standRepo "SleekSpace/repositories/property/stand"
	generalUtilities "SleekSpace/utilities/funcs/general"
	propertyUtilities "SleekSpace/utilities/funcs/property"

	"github.com/gin-gonic/gin"
)

func GetPropertyTypeById(c *gin.Context, propertyType string, propertyId int) {
	if propertyType == "stand" {
		stand := standRepo.GetStandWithAllAssociationsById(generalUtilities.ConvertIntToString(propertyId))
		if stand == nil {
			c.JSON(http.StatusForbidden, gin.H{"error": "this stand does not exist"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"response": propertyUtilities.PropertyStandResponse(*stand)})
		return
	}
}
