package routes

import (
	"todo-backend/controllers"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func AuthRoutes(r *gin.Engine, db *gorm.DB) {
	r.POST("/login", controllers.Login(db))
	r.POST("/signup", controllers.Signup(db))
	r.GET("/logout", controllers.Logout)
}
