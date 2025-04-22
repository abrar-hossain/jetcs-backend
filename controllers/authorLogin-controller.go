package controllers

import (
	"jetcs-backend/config"
	"jetcs-backend/db/migration/query"
	"jetcs-backend/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AuthorLogin(c *gin.Context) {
	var loginAuthor models.LoginInput

	if err := c.ShouldBindJSON(&loginAuthor); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Look up the author by email
	author, err := query.GetAuthorByEmail(config.DB, loginAuthor.Email)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"author": author})
}
