package repositories

import (
	constants "SleekSpace/utilities/constants"
	generalUtilities "SleekSpace/utilities/funcs/general"
	"math"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Paginate(c *gin.Context) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		pageStr := c.DefaultQuery("page", "1")
		page := generalUtilities.ConvertStringToInt(pageStr)
		if page <= 0 {
			page = 1
		}

		dbClone := db.Session(&gorm.Session{})
		var total int64
		dbClone.Count(&total)
		totalPages := int(math.Ceil(float64(total) / float64(constants.PageLimit)))

		c.Set("totalPages", totalPages)

		offSet := (page - 1) * constants.PageLimit
		return db.Offset(offSet).Limit(constants.PageLimit)
	}
}
