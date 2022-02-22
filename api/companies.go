package api

import (
	"strconv"
	"tvdapi/orm"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func HandleCompanyList(db *gorm.DB) gin.HandlerFunc {
	fn := func(c *gin.Context) {
		limit, err := strconv.Atoi(c.DefaultQuery("count", "10"))

		if err != nil {
			c.JSON(200, gin.H{"error": "Parameter count must be integer type."})
		} else {
			var companies []orm.Company

			// result := db.
			// 	Limit(limit).
			// 	Order("ID").
			// 	Find(&companies)

			result := db.
				Joins("Access", db.Where("Access.FiscalConfiguration = ?", true)).
				Limit(limit).
				Order("ID").
				Find(&companies)

			c.JSON(200, gin.H{
				"count": result.RowsAffected,
				"items": companies,
			})
		}
	}

	return gin.HandlerFunc(fn)
}
