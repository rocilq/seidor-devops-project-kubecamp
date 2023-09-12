package db

import (
	"context"
	"errors"
	"os"

	"github.com/gofiber/fiber/v2/log"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type DB struct {
	client *mongo.Client
}

func InitDB() *DB {
	db := DB{}
	if err := db.Connect(); err != nil {
		log.Error(err)
		panic("fail to connect to db")
	}
	return &db
}

func (db *DB) Connect() error {
	if db.client == nil {
		connectionUri := os.Getenv("MONGO_URI")
		if connectionUri == "" {
			return errors.New("could not find MONGO_URI env")
		}

		if client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(connectionUri)); err != nil {
			return err
		} else {
			db.client = client
			return nil
		}
	}

	return nil
}

func (db *DB) Close() error {
	if db.client == nil {
		return errors.New("the client that was trying to close dint exist")
	}

	return db.client.Disconnect(context.TODO())
}

func (db *DB) GetCollection(collectionName string, dbName string) (*mongo.Collection, error) {
	if db.client == nil {
		return nil, errors.New("unable to find the collection")
	}

	return db.client.Database(dbName).Collection(collectionName), nil
}

func (db *DB) UsersCollection() (*mongo.Collection, error) {
	if db.client == nil {
		return nil, errors.New("unable to find the collection")
	}

	return db.client.Database("auth-service").Collection("users"), nil
}
