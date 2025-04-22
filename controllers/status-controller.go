package controllers

import (
	"jetcs-backend/config"
	"jetcs-backend/db/migration/query"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CheckStatus(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	status, err := query.CheckSubmissionStatus(config.DB, id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"id":     id,
		"status": status,
	})
}
