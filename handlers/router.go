package handlers

import (
	"fmt"
	"log"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"

	"ginexample/middlewares/jwtauth"
	"ginexample/handlers/ping"
	"ginexample/middlewares/logs"
)

func SetupRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Recovery())
	r.Use(logs.CustomeLogMiddleware())

	authMiddleware := jwtauth.SetupAuthMiddleware()
	r.POST("/login", authMiddleware.LoginHandler)

	r.NoRoute(authMiddleware.MiddlewareFunc(), func(c *gin.Context) {
		claims := jwt.ExtractClaims(c)
		log.Printf("NoRoute claims: %#v\n", claims)
		c.JSON(404, gin.H{"message": "Page not found"})
	})

	auth := r.Group("/auth")
	// Refresh time can be longer than token timeout
	auth.GET("/refresh_token", authMiddleware.RefreshHandler)
	auth.Use(authMiddleware.MiddlewareFunc())
	{
		auth.GET("/hello", jwtauth.HelloHandler)
		auth.GET(ping.PingUri, ping.Ping)
	}
	return r
}

func HandleError(c *gin.Context, code int, err error) {
	msg := fmt.Sprintf("%s", err)
	c.AbortWithStatusJSON(code, gin.H{"msg": msg})
}
