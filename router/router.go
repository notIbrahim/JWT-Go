package router

import (
	"jwt-go/app/middleware"
	"jwt-go/app/query"

	"github.com/gin-gonic/gin"
)

func ReadyServer() *gin.Engine {
	route := gin.Default()

	UserRoute := route.Group("/users")
	{
		UserRoute.POST("/register", query.UserRegister)
		UserRoute.POST("/login", query.UserLogged)
	}

	RouteToProduct := route.Group("/products")
	{
		RouteToProduct.Use(middleware.Auth())
		RouteToProduct.POST("/", middleware.ProductAuthor(), query.CreateProduct)
		RouteToProduct.GET("/:ID", middleware.ProductAuthor(), query.ReadProduct)
		RouteToProduct.PUT("/:ID", middleware.ProductAuthor(), query.UpdateProduct)
		RouteToProduct.DELETE("/:ID", middleware.ProductAuthor(), query.DeleteProduct)
	}
	return route
}
