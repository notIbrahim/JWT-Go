package query

import (
	"jwt-go/app/entity"
	"jwt-go/pkg/database"
	"jwt-go/pkg/helper"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func ReadProduct(ResponseContext *gin.Context) {
	DB, _ := database.Connect()
	ReceivedContent := helper.GetContentTypeOf(ResponseContext)
	BaseProduct := entity.Product{}

	// Read Parametes by Convert Strconv All to Int
	ProductID, err := strconv.Atoi(ResponseContext.Param("ID"))
	if err != nil {
		ResponseContext.JSON(http.StatusBadRequest, gin.H{
			"error":   "Invaild Request",
			"Message": err.Error(),
		})
		return
	}
	// Check Content Type

	if ReceivedContent == app {
		ResponseContext.ShouldBindJSON(&BaseProduct)
	} else {
		ResponseContext.ShouldBind(&BaseProduct)
	}

	// Query
	err = DB.First(&BaseProduct, "ID = ?", ProductID).Error
	if err != nil {
		ResponseContext.JSON(http.StatusBadRequest, gin.H{
			"Error":   err.Error(),
			"Message": "Invalid Request ID",
		})
		return
	}

	ResponseContext.JSON(http.StatusOK, BaseProduct)
}

func CreateProduct(ResponseContext *gin.Context) {
	DB, _ := database.Connect()
	UserData := ResponseContext.MustGet("UserData").(jwt.MapClaims)
	BaseProduct := entity.Product{}
	ReceivedContent := helper.GetContentTypeOf(ResponseContext)

	UserID := uint(UserData["ID"].(float64))
	// Check Content Type

	if ReceivedContent == app {
		ResponseContext.ShouldBindJSON(&BaseProduct)
	} else {
		ResponseContext.ShouldBind(&BaseProduct)
	}

	BaseProduct.UserID = UserID
	err := DB.Create(&BaseProduct).Error

	if err != nil {
		ResponseContext.JSON(http.StatusBadRequest, gin.H{
			"Error":   err.Error(),
			"Message": "Request Invalid",
		})
		return
	}

	ResponseContext.JSON(http.StatusOK, BaseProduct)
}

func UpdateProduct(ResponseContext *gin.Context) {
	DB, _ := database.Connect()
	UserData := ResponseContext.MustGet("UserData").(jwt.MapClaims)
	BaseProduct := entity.Product{}
	ReceivedContent := helper.GetContentTypeOf(ResponseContext)
	UserLevels := UserData["Level"]

	if ReceivedContent == app {
		ResponseContext.ShouldBindJSON(&BaseProduct)
	} else {
		ResponseContext.ShouldBind(&BaseProduct)
	}

	if UserLevels == "user" {
		ResponseContext.JSON(http.StatusUnauthorized, gin.H{
			"Status":  http.StatusUnauthorized,
			"Message": "Unauthorized",
		})
		return
	}

	ProductID, _ := strconv.Atoi(ResponseContext.Param("ID"))
	UserID := uint(UserData["ID"].(float64))
	BaseProduct.ID = uint(ProductID)
	BaseProduct.UserID = UserID

	err := DB.Model(&BaseProduct).Where("ID = ?", ProductID).Updates(entity.Product{
		Title:       BaseProduct.Title,
		Description: BaseProduct.Description,
	}).Error

	if err != nil {
		ResponseContext.JSON(http.StatusBadRequest, gin.H{
			"Error":   err.Error(),
			"Message": "Request Invalid",
		})
		return
	}
	ResponseContext.JSON(http.StatusOK, BaseProduct)
}

func DeleteProduct(ResponseContext *gin.Context) {
	DB, _ := database.Connect()
	UserData := ResponseContext.MustGet("UserData").(jwt.MapClaims)
	BaseProduct := entity.Product{}
	ReceivedContent := helper.GetContentTypeOf(ResponseContext)
	UserLevels := UserData["Level"]

	ProductID, _ := strconv.Atoi(ResponseContext.Param("ID"))
	if ReceivedContent == app {
		ResponseContext.ShouldBindJSON(&BaseProduct)
	} else {
		ResponseContext.ShouldBind(&BaseProduct)
	}

	if UserLevels == "user" {
		ResponseContext.JSON(http.StatusUnauthorized, gin.H{
			"Status":  http.StatusUnauthorized,
			"Message": "Unauthorized",
		})
		return
	}

	err := DB.Where("ID = ?", ProductID).Delete(&BaseProduct).Error
	if err != nil {
		ResponseContext.JSON(http.StatusBadRequest, gin.H{
			"Error":   err.Error(),
			"Message": "Request Invalid",
		})
		return
	}
	ResponseContext.JSON(http.StatusAccepted, gin.H{
		"Status":  http.StatusAccepted,
		"Message": "Data Berhasil di hapus",
	})
}
