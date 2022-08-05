package controllers

import (
	"gorm-gin-practise/initializers"
	"gorm-gin-practise/models"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

type signUpSerializer struct {
	Email     string `json:"email"`
	Password  string `json:"password"`
	Password2 string `json:"password2"`
}

type loginSerializer struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// Sign Up godoc
// @Summary      Sign Up
// @Description  Sign Up user
// @Tags         Auth
// @Accept       json
// @Produce      json
// @Param        user body signUpSerializer  true  "Sign Up"
// @Success      201  {string}  http.StatusOK
// @Router       /auth/siginup [post]
func SignUp(c *gin.Context) {
	// Get the email/password off request body
	var body signUpSerializer
	if c.ShouldBindJSON(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Faild to read body",
		})
		return
	}
	if body.Password != body.Password2 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "passwords are not equal each other",
		})
		return
	}
	// Hash the password
	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), 10)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Faild to hash password",
		})
		return
	}
	// Create the user
	tx := initializers.DB.Create(&models.User{Email: body.Email, Password: string(hash)})
	if tx.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Faild to create user",
		})
		return
	}
	// Respond
	c.JSON(http.StatusCreated, gin.H{"message": "Created User"})
}

// Login godoc
// @Summary      Login
// @Description  Login user
// @Tags         Auth
// @Accept       json
// @Produce      json
// @Param        user body loginSerializer  true  "Login"
// @Success      200  {string}  http.StatusOK
// @Router       /auth/login [post]
func Login(c *gin.Context) {
	// Get the email/password off request body
	var body loginSerializer
	if c.ShouldBindJSON(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Faild to read body",
		})
		return
	}
	// Look up requested user
	var user models.User
	initializers.DB.First(&user, "email = ?", body.Email)
	if user.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid email or password",
		})
		return
	}
	// Compare sent in pass with saved user pass hash
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password)); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid email or password",
		})
		return
	}
	// Generate a jwt token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(time.Hour * 3).Unix(),
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString([]byte(os.Getenv("SECERET")))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Faild to create token",
		})
		return
	}
	// Send it back
	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorization", tokenString, 3600*3, "", "", false, true)

	// c.JSON(http.StatusOK, gin.H{"token": tokenString})
	c.JSON(http.StatusOK, gin.H{})
}
