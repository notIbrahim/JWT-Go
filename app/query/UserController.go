package query

import (
	"jwt-go/app/entity"
	"jwt-go/pkg/database"
	"jwt-go/pkg/helper"
	"net/http"

	"github.com/gin-gonic/gin"
)

const app = "application/json"

func UserRegister(ResponseContext *gin.Context) {
	DB, err := database.Connect()
	if err != nil {
		panic(err)
	}
	ReceivedContent := helper.GetContentTypeOf(ResponseContext)
	_, _ = DB, ReceivedContent
	Users := entity.Users{}

	if ReceivedContent == app {
		ResponseContext.ShouldBindJSON(&Users)
	} else {
		ResponseContext.ShouldBind(&Users)
	}

	err = DB.Create(&Users).Error
	if err != nil {
		ResponseContext.JSON(http.StatusBadRequest, gin.H{
			"Error":   err.Error(),
			"Message": "Invalid Create User request",
		})
		return
	}

	ResponseContext.JSON(http.StatusAccepted, gin.H{
		"ID":       Users.ID,
		"email":    Users.Email,
		"Fullname": Users.Fullname,
		"level":    Users.Level,
	})

}

func UserLogged(ResponseContext *gin.Context) {
	DB, err := database.Connect()
	if err != nil {
		panic(err)
	}
	ReceivedContent := helper.GetContentTypeOf(ResponseContext)
	_, _ = DB, ReceivedContent
	Users := entity.Users{}
	Password := ""

	if ReceivedContent == app {
		ResponseContext.ShouldBindJSON(&Users)
	} else {
		ResponseContext.ShouldBind(&Users)
	}

	SeizePassword := helper.PasswordCheck([]byte(Users.Password), []byte(Password))

	if !SeizePassword {
		ResponseContext.JSON(http.StatusUnauthorized, gin.H{
			"Status":  http.StatusUnauthorized,
			"Message": "Invalid Password or Email",
		})
		return
	}
	GenerateToken := helper.GenKeys(Users.ID, Users.Email, Users.Level)
	ResponseContext.JSON(http.StatusAccepted, gin.H{
		"Tokens": GenerateToken,
	})

}
