package middleware

import (
	"jwt-go/app/entity"
	"jwt-go/pkg/database"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func ProductAuthor() gin.HandlerFunc {
	return func(ReceivedContext *gin.Context) {
		DB, _ := database.Connect()
		ProductID, err := strconv.Atoi(ReceivedContext.Param("ID"))
		if err != nil {
			ReceivedContext.JSON(http.StatusBadRequest, gin.H{
				"Error":   http.StatusBadRequest,
				"Message": "Missing ID Parameter or Invalid",
			})
			return
		}

		UserData := ReceivedContext.MustGet("UserData").(jwt.MapClaims)
		UserID := uint(UserData["ID"].(float64))
		UserLevel := UserData["Level"]
		Product := entity.Product{}
		// If i correctly called they finding user_id by referencing Pointer Location Product
		// and then Parse it using unsigned integer
		err = DB.Select("user_id").First(&Product, uint(ProductID)).Error
		if err != nil {
			ReceivedContext.JSON(http.StatusNotFound, gin.H{
				"Error":   http.StatusNotFound,
				"Message": "Either Data Doesn't exist or Missing Somewhere Else",
			})
			return
		}
		// Finding Similiarity between UserID and UserLevel
		if Product.UserID != UserID && UserLevel == "user" {
			ReceivedContext.JSON(http.StatusUnauthorized, gin.H{
				"Error":   http.StatusUnauthorized,
				"Message": "Unallowed to find this data",
			})
		}
	}
}
