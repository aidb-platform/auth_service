package controllers

import (
	"net/http"

	"github.com/aidb-platform/auth_service/models"
	"github.com/aidb-platform/auth_service/utils"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type SignupRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
	Name     string `json:"name"`
}

func Signup(c *gin.Context) {
	var body SignupRequest
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	hashedPassword, err := utils.HashPassword(body.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to hash password"})
		return
	}

	// Auto-create org for now (can be improved later)
	org := models.Organization{
		ID:   uuid.New(),
		Name: body.Name + "'s Org",
	}
	if err := models.DB.Create(&org).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not create organization"})
		return
	}

	user := models.User{
		ID:           uuid.New(),
		Email:        body.Email,
		Name:         body.Name,
		PasswordHash: hashedPassword,
		OrgID:        org.ID,
		IsVerified:   false,
		IsAdmin: false,
	}

	if err := models.DB.Create(&user).Error; err != nil {
		c.JSON(http.StatusConflict, gin.H{"error": "User already exists or invalid"})
		return
	}

	token, _ := utils.GenerateToken(user.ID.String(), org.ID.String())
	c.JSON(http.StatusOK, gin.H{
		"token": token,
		"user":  user.Email,
		"org":   org.Name,
	})
}

type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

func Login(c *gin.Context) {
	var body LoginRequest
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var user models.User
	if err := models.DB.Where("email = ?", body.Email).First(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
		return
	}

	if !utils.CheckPasswordHash(body.Password, user.PasswordHash) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
		return
	}

	token, _ := utils.GenerateToken(user.ID.String(), user.OrgID.String())
	c.JSON(http.StatusOK, gin.H{
		"token": token,
		"user":  user.Email,
	})
}
