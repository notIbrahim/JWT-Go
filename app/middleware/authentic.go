package middleware

import (
	"github.com/gin-gonic/gin"
	"jwt-go/pkg/helper"
	"net/http"
)

func Auth() gin.HandlerFunc {
	return func(Context *gin.Context) {
		ReceivedToken, err := helper.VerficationTokenizer(Context)
		_ = ReceivedToken
		if err != nil {
			Context.JSON(http.StatusUnauthorized, gin.H{
				"Status":  http.StatusUnauthorized,
				"Message": err.Error(),
			})
			return
		}
		Context.Set("UserData", ReceivedToken)
		Context.Next()
	}
}
