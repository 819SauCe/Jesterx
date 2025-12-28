package services

import (
	"database/sql"
	"jesterx-core/config"
	"jesterx-core/helpers"
	"jesterx-core/responses"

	"github.com/gin-gonic/gin"
)

type applyThemeBody struct {
	ThemeID string `json:"theme_id"`
}

func ApplyThemeService(c *gin.Context) {
	user := c.MustGet("user").(helpers.UserData)

	tenantSlug := c.GetHeader("X-Tenant-Page-Id")
	if tenantSlug == "" {
		c.JSON(400, responses.ErrorResponse{Success: false, Message: "Missing tenant header."})
		return
	}

	var body applyThemeBody
	if err := c.ShouldBindJSON(&body); err != nil || body.ThemeID == "" {
		c.JSON(400, responses.ErrorResponse{Success: false, Message: "Invalid theme_id."})
		return
	}

	db := config.DB

	var tenantID string
	err := db.QueryRow(`SELECT id FROM tenants WHERE page_id = $1`, tenantSlug).Scan(&tenantID)
	if err == sql.ErrNoRows {
		c.JSON(404, responses.ErrorResponse{Success: false, Message: "Tenant not found."})
		return
	}
	if err != nil {
		c.JSON(500, responses.ErrorResponse{Success: false, Message: "Failed to resolve tenant."})
		return
	}

	hasAccess, _, err := helpers.UserHasTenantAccess(user.Id, tenantID)
	if err != nil {
		c.JSON(500, responses.ErrorResponse{Success: false, Message: "Failed to check permissions."})
		return
	}
	if !hasAccess {
		c.JSON(403, responses.ErrorResponse{Success: false, Message: "You do not belong to this site (tenant)."})
		return
	}

	_, err = db.Exec(`UPDATE pages SET theme_id = $1, updated_at = NOW() WHERE tenant_id = $2`, body.ThemeID, tenantID)
	if err != nil {
		c.JSON(500, responses.ErrorResponse{Success: false, Message: "Failed to apply theme."})
		return
	}

	c.JSON(200, gin.H{
		"success": true,
		"message": "",
		"data": gin.H{
			"theme_id": body.ThemeID,
		},
	})
}
