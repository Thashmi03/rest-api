package routes

import (
	"rest-api/controllers"

	"github.com/gin-gonic/gin"
)

func TransactionRoutes(router *gin.Engine,controller controllers.TransactionController){
	router.POST("/api/create",controller.CreateTransaction)
	router.GET("/api/get",controller.GetAllTransaction)
}