package mongodb

import (
	"context"
	"fmt"
	"strings"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type history struct {
	Name string `bson:"name"`
	Text string `bson:"text"`
}

func (mdb *LogsDB) PingDB(ctx context.Context) error {
	if err := mdb.client.Ping(ctx, nil); err != nil {
		return err
	}
	return nil
}

func (mdb *LogsDB) StoreToChatLogDB(ctx context.Context, name string, text string) error {
	fmt.Println("Name: ", name, " Text: ", text)
	hist := history{Name: name, Text: text}
	_, err := mdb.collection.InsertOne(ctx, hist)
	if err != nil {
		return err
	}
	return nil
}

func (mdb *LogsDB) ReadFromChatLogDB(ctx context.Context, val int64) (string, error) {
	if val == 0 {
		val = 200
	}

	opts := options.Find().
		SetSort(bson.D{{Key: "_id", Value: -1}}).SetLimit(val)

	res, err := mdb.collection.Find(ctx, bson.M{}, opts)
	if err != nil {
		return "", err
	}
	defer res.Close(ctx)

	var his []history
	if err := res.All(ctx, &his); err != nil {
		return "", err
	}

	lenHis := len(his)

	var sb strings.Builder
	sb.Grow(lenHis)

	for _, his := range his {
		sb.WriteString(fmt.Sprintf("Name: %s, Text:%s", his.Name, his.Text))
	}
	return sb.String(), nil
}
