package handlers

import (
	"ginexample/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {
	version := models.GetMysqlVersion()
	c.JSON(http.StatusOK, gin.H{"message": version})
}
