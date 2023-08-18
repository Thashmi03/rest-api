package controllers

import (
	"net/http"
	"rest-api/interfaces"
	"rest-api/models"
	"strings"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TransactionController struct{
	trans interfaces.ITransaction
}
func InitController(profile interfaces.ITransaction)TransactionController{
	return TransactionController{profile}
}
func (t * TransactionController)CreateTransaction(ctx *gin.Context){
	var transfer *models.Transaction
	if err := ctx.ShouldBindJSON(&transfer); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}
	newProfile, err := t.trans.CreateTransactions(transfer)

	if err != nil {
		if strings.Contains(err.Error(), "title already exists") {
			ctx.JSON(http.StatusConflict, gin.H{"status": "fail", "message": err.Error()})
			return
		}

		ctx.JSON(http.StatusBadGateway, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"status": "success", "data": newProfile})
}
func  (t * TransactionController)GetAllTransaction(ctx *gin.Context){
	id,_:= primitive.ObjectIDFromHex("64df4309ff5af24396f2bb8a")
	newProfile, err := t.trans.GetTransactionByCustomerId(id)

	if err != nil {
		// if strings.Contains(err.Error(), "title already exists") {
		// 	ctx.JSON(http.StatusConflict, gin.H{"status": "fail", "message": err.Error()})
		// 	return
		// }

		ctx.JSON(http.StatusBadGateway, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"status": "success", "data": newProfile})
	ctx.String(http.StatusOK,"got the data")
}