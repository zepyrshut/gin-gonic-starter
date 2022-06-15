package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (m *Repository) GetOneUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status": "malformed id",
			"error":  err.Error(),
		})
		return
	}

	user, err := m.DB.OneUser(id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"status": "error",
			"error":  err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
		"user":   user,
	})
}
