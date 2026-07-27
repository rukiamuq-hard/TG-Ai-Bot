package mongodb

import (
	"context"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type LogsDB struct {
	client     *mongo.Client
	collection *mongo.Collection
}

func New() *LogsDB {
	return &LogsDB{}
}

func (mdb *LogsDB) ConnectDB() error {

	var err error
	mdb.client, err = mongo.Connect(options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		return err
	}
	mdb.collection = mdb.client.Database("Logs").Collection("History")

	if err := mdb.PingDB(context.Background()); err != nil {
		return err
	}

	return nil
}

func (mdb *LogsDB) Close() {
	mdb.client.Disconnect(context.Background())
}
