package interfaces

import (
	"rest-api/models"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type ITransaction interface{
	CreateTransactions(profile *models.Transaction)(*mongo.InsertOneResult,error)
	GetTransactionByCustomerId(profile primitive.ObjectID)([]*models.Transaction,error)
}