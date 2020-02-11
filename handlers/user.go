package handlers

import (
	"net/http"
	"ginexample/models"
	"ginexample/common/mysql"

	"github.com/gin-gonic/gin"
)

type userRequest struct {
	Name string `json:"name" binding:"required,max=100,min=1"`
}

func InsertUser(c *gin.Context) {
	userReq := new(userRequest)
	if err := c.ShouldBindJSON(userReq); err != nil {
		HandleError(c, http.StatusBadRequest, err)
		return
	}
	user := models.User{Name: userReq.Name}

	count, err := user.InsertUser(mysql.SQLX_DB)
	if err != nil{
		HandleError(c, http.StatusInternalServerError, err)
	}
	c.JSON(http.StatusOK,gin.H{"count":count})
}
