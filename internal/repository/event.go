package repository

import (
	"context"

	"github.com/audi-skripsi/lambda_event_presenter/internal/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (r *repository) InsertEvent(event model.EventLog, collName string) (result model.EventLog, err error) {
	coll := r.mongo.Collection(collName)

	res, err := coll.InsertOne(context.Background(), event)
	if err != nil {
		r.logger.Errorf("error inserting to mongodb for %s: %+v", event.UID, err)
		return
	}

	objID, ok := res.InsertedID.(primitive.ObjectID)
	if !ok {
		r.logger.Errorf("error asserting object id for %s: %+v", event.UID, err)
		return
	}

	event.ObjectID = objID
	result = event
	return
}
