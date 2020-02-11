package handlers

import (
	"fmt"

	"github.com/gin-gonic/gin"

	"ginexample/middlewares/logs"
)

func SetupRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Recovery())

	g := r.Group("/api")
	g.Use(logs.CustomeLogMiddleware())
	{
		g.GET("/index", Index)
		g.POST("/user", InsertUser)
	}
	return r
}

func HandleError(c *gin.Context, code int, err error) {
	msg := fmt.Sprintf("%s", err)
	c.AbortWithStatusJSON(code, gin.H{"msg": msg})
}
