package services

import (
	"jesterx-core/config"
	"jesterx-core/helpers"
	"jesterx-core/responses"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

func DeletePageService(c *gin.Context) {
	user := c.MustGet("user").(helpers.UserData)
	tenantID := c.MustGet("tenantID").(string)
	pageID := c.Param("page_id")

	if pageID == "" {
		c.JSON(400, responses.ErrorResponse{Success: false, Message: "Invalid page id."})
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

	db := config.DB

	result, err := db.Exec(`DELETE FROM pages WHERE tenant_id = $1 AND page_id = $2`, tenantID, pageID)
	if err != nil {
		c.JSON(500, responses.ErrorResponse{Success: false, Message: "Failed to delete page in Postgres."})
		return
	}

	affected, err := result.RowsAffected()
	if err != nil {
		c.JSON(500, responses.ErrorResponse{Success: false, Message: "Failed to delete page in Postgres."})
		return
	}
	if affected == 0 {
		c.JSON(404, responses.ErrorResponse{Success: false, Message: "Page not found."})
		return
	}

	collection := config.MongoClient.Database("genyou").Collection("page_sveltes")
	_, _ = collection.DeleteOne(c.Request.Context(), bson.M{
		"tenant_id": tenantID,
		"page_id":   pageID,
	})

	c.Status(204)
}
