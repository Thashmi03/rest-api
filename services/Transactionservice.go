package services

import (
	"context"
	"errors"
	"rest-api/interfaces"
	"rest-api/models"
	"time"

	// "time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	// "go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Context struct {
	Collection *mongo.Collection
	ctx        context.Context
}

func InitializeTransactionService(collection *mongo.Collection, ctx context.Context) interfaces.ITransaction {
	return &Context{collection, ctx}
}
func (s *Context) CreateTransactions(user *models.Transaction) (*mongo.InsertOneResult, error) {
	user.Id = primitive.NewObjectID()
	user.TransactionDate = time.Now()

	result, err := s.Collection.InsertOne(s.ctx, &user)
	if err != nil {
		if er, ok := err.(mongo.WriteException); ok && er.WriteErrors[0].Code == 11000 {
			return nil, errors.New("user with that email already exist")
		}
		return nil, err
	}
	return result, err
}
func (s *Context) GetTransactionByCustomerId(id primitive.ObjectID) ([]*models.Transaction, error) {
	// filter:=bson.D{}
	matchStage := bson.D{{Key: "_id", Value: id}}
	result, err := s.Collection.Find(s.ctx, matchStage)
	if err != nil {
		return nil, err

	} else {
		var transaction []*models.Transaction
		for result.Next(s.ctx) {
			transfer := &models.Transaction{}
			err := result.Decode(transfer)
			if err != nil {
				return nil, err
			}
			transaction = append(transaction, transfer)
		}
		return transaction, nil
	}
}
