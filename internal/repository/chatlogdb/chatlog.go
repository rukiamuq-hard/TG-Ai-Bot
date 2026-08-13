package mongodb

import (
	"TgAiBot/internal/apperrors"
	"TgAiBot/internal/models"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"strings"
)

func (mdb *LogsDB) StoreToChatLogDB(ctx context.Context, hist models.History) error {
	_, err := mdb.collection.InsertOne(ctx, hist)
	if err != nil {
		return err
	}
	return nil
}

func (mdb *LogsDB) ReadFromChatLogDB(ctx context.Context, val int64) ([]models.History, string, error) {
	opts := options.Find().
		SetSort(bson.D{{Key: "_id", Value: -1}}).SetLimit(val)

	res, err := mdb.collection.Find(ctx, bson.M{}, opts)
	if err != nil {
		return nil, "", err
	}
	defer res.Close(ctx)

	var his []models.History
	if err := res.All(ctx, &his); err != nil {
		return nil, "", err
	}

	lenHis := len(his)

	var sb strings.Builder
	sb.Grow(lenHis)

	for _, his := range his {
		logs := fmt.Sprintf("Name: %s, UID:%d, Text:%s, CID:%d, MID:%d", his.Name, his.UID, his.Text, his.CID, his.MID)
		sb.WriteString(logs)
	}
	return his, sb.String(), nil
}

func (mdb *LogsDB) DeleteMessage(ctx context.Context, chat_id int64, val int64) error {
	options := options.Find().
		SetSort(bson.D{
			{Key: "mid", Value: -1},
		}).
		SetLimit(val)

	cursor, err := mdb.collection.Find(ctx, bson.M{
		"cid": chat_id,
	}, options)

	if err != nil {
		return err
	}
	defer cursor.Close(ctx)

	var docs []struct {
		ID any `bson:"_id"`
	}

	if err = cursor.All(ctx, &docs); err != nil {
		return err
	}

	if len(docs) == 0 {
		return apperrors.NoMessageToDelete
	}

	ids := make([]any, len(docs))
	for i, d := range docs {
		ids[i] = d.ID
	}

	_, err = mdb.collection.DeleteMany(ctx, bson.M{
		"_id": bson.M{
			"$in": ids,
		},
	})

	if err != nil {
		return err
	}

	return nil
}
