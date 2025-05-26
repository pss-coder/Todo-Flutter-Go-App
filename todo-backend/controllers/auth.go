package controllers

import (
	"net/http"
	"time"
	"todo-backend/models"
	"todo-backend/utils"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// TODO: USE PRIVATE KEY FOR JWT SIGNING
var jwtKey = []byte("my_secret_key")

// Login godoc
// @Summary      Login user
// @Description  Authenticates user and sets a JWT token in a cookie
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        user  body      models.User  true  "User credentials"
// @Success      200  {object}  map[string]interface{}
// @Failure      400  {object}  map[string]string
// @Failure      401  {object}  map[string]string
// @Router       /login [post]
func Login(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var user models.User

		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		var existingUser models.User

		db.Where("email = ?", user.Email).First(&existingUser)

		if existingUser.ID == 0 {
			c.JSON(401, gin.H{"error": "user does not exist"})
			return
		}

		errHash := utils.CompareHashPassword(user.Password, existingUser.Password)

		if !errHash {
			c.JSON(401, gin.H{"error": "Invalid password"})
			return
		}

		expirationTime := time.Now().Add(5 * time.Minute)

		claims := &models.Claims{
			Role: existingUser.Role,
			StandardClaims: jwt.StandardClaims{
				Subject:   existingUser.Email,
				ExpiresAt: expirationTime.Unix(),
			},
		}

		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

		tokenString, err := token.SignedString(jwtKey)

		if err != nil {
			c.JSON(500, gin.H{"error": "Could not generate token"})
			return
		}

		c.SetCookie("token", tokenString, int(expirationTime.Unix()), "/", "localhost", false, true)
		// c.JSON(200, gin.H{"success": "user logged in"})
		c.JSON(http.StatusOK, struct {
			Name  string `json:"name"`
			Email string `json:"email"`
			Role  string `json:"role"`
		}{
			Name:  existingUser.Name,
			Email: existingUser.Email,
			Role:  existingUser.Role,
		})

	}
}

// Signup godoc
// @Summary      Signup user
// @Description  Registers a new user
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        user  body      models.User  true  "User info"
// @Success      200  {object}  map[string]string
// @Failure      400  {object}  map[string]string
// @Router       /signup [post]
func Signup(db *gorm.DB) gin.HandlerFunc {

	return func(c *gin.Context) {
		var user models.User

		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		var existingUser models.User

		db.Where("email = ?", user.Email).First(&existingUser)

		if existingUser.ID != 0 {
			c.JSON(400, gin.H{"error": "user already exists"})
			return
		}

		var errHash error
		user.Password, errHash = utils.GenerateHashPassword(user.Password)

		if errHash != nil {
			c.JSON(500, gin.H{"error": "could not generate password hash"})
			return
		}

		db.Create(&user)
		c.JSON(200, gin.H{"success": "user created"})
	}
}

func Home(c *gin.Context) {
	cookie, err := c.Cookie("token")

	if err != nil {
		c.JSON(401, gin.H{"error": "unauthorized"})
		return
	}

	claims, err := utils.ParseToken(cookie)

	if err != nil {
		c.JSON(401, gin.H{"error": "unauthorized"})
		return
	}

	if claims.Role != "user" && claims.Role != "admin" {
		c.JSON(403, gin.H{"error": "unauthorized"})
		return
	}

	c.JSON(200, gin.H{"success": "welcome to the home page", "role": claims.Role})
}

func Premium(c *gin.Context) {
	cookie, err := c.Cookie("token")

	if err != nil {
		c.JSON(401, gin.H{"error": "unauthorized"})
		return
	}

	claims, err := utils.ParseToken(cookie)

	if err != nil {
		c.JSON(401, gin.H{"error": "unauthorized"})
		return
	}

	if claims.Role != "admin" {
		c.JSON(403, gin.H{"error": "unauthorized"})
		return
	}

	c.JSON(200, gin.H{"success": "welcome to the premium page", "role": claims.Role})
}

// Logout godoc
// @Summary      Logout user
// @Description  Clears the authentication cookie
// @Tags         auth
// @Produce      json
// @Success      200  {object}  map[string]string
// @Router       /logout [get]
func Logout(c *gin.Context) {
	c.SetCookie("token", "", -1, "/", "localhost", false, true)
	c.JSON(200, gin.H{"success": "user logged out"})
}
