package server

import (
	"jetcs-backend/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
	r.POST("/signup", controllers.CreateAuthor)
	r.POST("/login", controllers.AuthorLogin)
	r.POST("/submit", controllers.SubmitPaper)
	r.GET("/status/:id", controllers.CheckStatus)
}
