package routes

import (
	"net/http"

	"github.com/aidb-platform/auth_service/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// CurrentUser handles GET /api/me
func CurrentUser(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get user ID from middleware context
		userIDValue, exists := c.Get("userID")
		if !exists {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized - user ID missing"})
			return
		}

		userID, ok := userIDValue.(string)
		if !ok {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "User ID type assertion failed"})
			return
		}

		var user models.User
		if err := db.First(&user, "id = ?", userID).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"success": true,
			"user": gin.H{
				"id":        user.ID,
				"email":     user.Email,
				"org_id":    user.OrgID,
				"is_admin":  user.IsAdmin,
				"createdAt": user.CreatedAt,
			},
		})
	}
}
