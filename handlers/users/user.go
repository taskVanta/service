package handlers

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"service/config"
	models "service/models/users"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

// Signup handler
func Signup(c *gin.Context) {
	var input models.User
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Check if a user with this email exists
	var existingUser models.User
	err := config.DB.Where("email = ?", input.Email).First(&existingUser).Error
	if err == nil {
		// User already exists
		c.JSON(http.StatusBadRequest, gin.H{"error": "Email already registered"})
		return
	} else if err != gorm.ErrRecordNotFound {
		// Unexpected DB error
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
		return
	}

	// Hash password and create user as before
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.PasswordHash), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Password hashing failed"})
		return
	}
	// Debugging line
	user := models.User{
		FirstName:    input.FirstName,
		Email:        input.Email,
		PasswordHash: string(hashedPassword),
	}

	if err := config.DB.Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "User creation failed"})
		return
	}

	token, err := generateJWT(user.Email)
	if err != nil {
		log.Printf("JWT generation error: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Token generation failed"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token, "user": user})
}

// Signin handler
func Signin(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var tempuser models.User
	if err := config.DB.Where("email = ?", user.Email).First(&tempuser).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(tempuser.PasswordHash), []byte(user.PasswordHash)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	token, err := generateJWT(tempuser.Email)

	if err != nil {
		log.Printf("JWT generation error: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Token generation failed"})
		return
	}
	userResponse := struct {
		ID        string `json:"id"`
		Email     string `json:"email"`
		FirstName string `json:"first_name,omitempty"`
		LastName  string `json:"last_name,omitempty"`
		Role      string `json:"role,omitempty"`
	}{
		ID:        tempuser.ID,
		Email:     tempuser.Email,
		FirstName: tempuser.FirstName,
		LastName:  tempuser.LastName,
		Role:      tempuser.Role,
	}

	c.JSON(http.StatusOK, gin.H{"user": userResponse, "token": token})
}

// JWT generation example
func generateJWT(email string) (string, error) {
	// Define token claims
	claims := jwt.MapClaims{
		"email": email,
		"exp":   time.Now().Add(time.Minute * 5).Unix(), // token expires in 72 hours
		"iat":   time.Now().Unix(),
	}

	// Create token with claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Get secret key from environment variable or config
	secret := []byte(os.Getenv("JWT_SECRET"))
	if len(secret) == 0 {
		return "", fmt.Errorf("JWT secret is not set")
	}

	// Sign and get the complete encoded token as a string
	tokenString, err := token.SignedString(secret)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func CreateUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// Set CreatedAt to current time
	user.CreatedAt = time.Now().Unix()
	if err := config.DB.Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, user)
}

func GetUserByID(c *gin.Context) {
	// Parse user ID from the URL parameter
	id := c.Param("id")

	// Find the user record by ID
	var user models.User
	if err := config.DB.First(&user, "id = ?", id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	fmt.Println("Retrieved user: %+v\n", user)

	// Return success response
	c.JSON(http.StatusOK, user)
}

func GetUsers(c *gin.Context) {
	var users []models.User
	config.DB.Find(&users)
	c.JSON(http.StatusOK, users)
}

func UpdateUser(c *gin.Context) {
	// Get ID from URL param
	id := c.Param("id")

	// Find existing user
	var user models.User
	if err := config.DB.First(&user, "id = ?", id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	// Bind JSON payload to user struct (partial update)
	var input struct {
		FirstName *string  `json:"name"`
		LastName  *string  `json:"last_name"`
		Username  **string `json:"username"`
		Role      *string  `json:"role"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Update fields only if provided in input (i.e., partial update)
	if input.FirstName != nil {
		user.FirstName = *input.FirstName
	}

	if input.LastName != nil {
		user.LastName = *input.LastName
	}
	if input.Username != nil {
		user.Username = *input.Username
	}
	if input.Role != nil {
		user.Role = *input.Role
	}

	user.UpdatedAt = time.Now().Unix()
	// Save the updated user
	if err := config.DB.Save(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User updated successfully", "user": user})
}

func DeleteUser(c *gin.Context) {
	// Parse user ID from the URL parameter

	id := c.Param("id")
	fmt.Println("Deleting user with ID:", id)
	// Find the user record by ID
	var user models.User
	if err := config.DB.First(&user, "id = ?", id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	// Delete the user record
	if err := config.DB.Delete(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete user"})
		return
	}

	// Return success response
	c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
}
