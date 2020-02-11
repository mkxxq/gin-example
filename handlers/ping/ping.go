package ping

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"ginexample/models"
)

const (
	PingUri = "/ping"
)

func pingRouter() *gin.Engine {
	router := gin.New()
	router.Use(gin.Recovery())
	authorized := router.Group("/")
	{
		authorized.GET(PingUri, Ping)
	}
	return router
}

func Ping(c *gin.Context) {
	version := models.GetMysqlVersion()
	c.JSON(http.StatusOK, gin.H{"message": version})
}
