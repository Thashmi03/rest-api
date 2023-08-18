package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Transaction struct{
	Id primitive.ObjectID `json:"id" bson:"_id"`
	TransactionDate time.Time`json:"transaction_date" bson:"transaction_date"`
	Amount int `json:"amount" bson:"amount"`
	TransactionType string `json:"transaction_type" bson:"transaction_type"`
	CustomerId int`json:"customer_id" bson:"customer_id"`
}