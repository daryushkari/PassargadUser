package rest

import (
	"PassargadUser/entities/requestModel"
	"github.com/gin-gonic/gin"
	"net/http"
)

func DeleteUser(c *gin.Context) {

}

func CreateUser(c *gin.Context) {
	var input requestModel.CreateRequest
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

}

func UpdateUser(c *gin.Context) {

}

func GetUserInfo(c *gin.Context) {

}

func LoginUser(c *gin.Context) {

}
