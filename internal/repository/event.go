package repository

import (
	"context"

	"github.com/audi-skripsi/lambda_event_presenter/internal/dto"
	"github.com/audi-skripsi/lambda_event_presenter/internal/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (r *repository) InsertEvent(event model.EventLog, collName string) (result model.EventLog, err error) {
	coll := r.mongo.Collection(collName)

	updateOptions := options.Update().SetUpsert(true)

	res, err := coll.UpdateOne(context.Background(),
		bson.M{
			"uid": event.UID,
		},
		bson.M{
			"$set": event,
		},
		updateOptions,
	)

	if err != nil {
		r.logger.Errorf("error inserting to mongodb for %s: %+v", event.UID, err)
		return
	}

	if res.MatchedCount == 0 {
		objID, ok := res.UpsertedID.(primitive.ObjectID)
		if !ok {
			r.logger.Errorf("error asserting object id for %s: %+v", event.UID, err)
			return
		}
		event.ObjectID = objID
	}

	result = event
	return
}

func (r *repository) BatchInsertEvent(eventBatch *dto.EventBatch) (err error) {
	coll := r.mongo.Collection(eventBatch.CollName)

	var models []mongo.WriteModel

	for _, v := range eventBatch.EventData {
		models = append(models,
			mongo.NewReplaceOneModel().SetFilter(
				bson.M{"uid": v.UID},
			).SetReplacement(
				v,
			).SetUpsert(true),
		)
	}

	_, err = coll.BulkWrite(context.Background(), models)
	if err != nil {
		r.logger.Errorf("error bulk write: %+v", err)
	}
	return
}
