package services

import (
	"database/sql"
	"gen-you-ecommerce/config"
	"gen-you-ecommerce/helpers"

	"github.com/gin-gonic/gin"
)

func LogoutService(c *gin.Context) {
	tenantPageID := c.GetHeader("X-Tenant-Page-Id")
	if tenantPageID == "" {
		c.JSON(400, gin.H{"success": false, "error": "Tenant not informed."})
		return
	}

	var tenantID string
	err := config.DB.QueryRow(`SELECT id FROM tenants WHERE page_id = $1`, tenantPageID).Scan(&tenantID)

	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(404, gin.H{"success": false, "error": "Tenant not found"})
			return
		}
		c.JSON(500, gin.H{"success": false, "error": "Database error"})
		return
	}

	helpers.SetAuthCookie(c, "", 0)
	c.JSON(200, "sucess")
}
