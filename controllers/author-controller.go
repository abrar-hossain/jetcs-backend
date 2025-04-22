package controllers

import (
	"fmt"
	"jetcs-backend/config"
	"jetcs-backend/db/migration/query"
	"jetcs-backend/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateAuthor(c *gin.Context) {
	var author models.Author

	if err := c.ShouldBindJSON(&author); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Print input for debugging
	fmt.Println("Registering:", author)

	// Save to DB
	err := query.SaveAuthor(config.DB, &author)
	if err != nil {
		// Print the real error in Postman and terminal
		fmt.Println("DB Error:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, author)
}
