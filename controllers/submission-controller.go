package controllers

import (
	"jetcs-backend/config"
	"jetcs-backend/db/migration/query"
	"jetcs-backend/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SubmitPaper(c *gin.Context) {
	var submission models.Submission

	if err := c.ShouldBindJSON(&submission); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := query.CreateSubmission(config.DB, &submission)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to submit paper"})
		return
	}
	c.JSON(http.StatusOK, submission)

}
