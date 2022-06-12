package repository

import (
	"context"
	"fmt"

	indto "github.com/audi-skripsi/lambda_event_presenter/internal/dto"
	"github.com/audi-skripsi/lambda_event_presenter/internal/model"
	"github.com/audi-skripsi/lambda_event_presenter/internal/util/microserviceutil"
	"github.com/audi-skripsi/lambda_event_presenter/pkg/dto"

	"github.com/audi-skripsi/lambda_event_presenter/pkg/errors"
	"github.com/audi-skripsi/lambda_event_presenter/pkg/util/jsonutil"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	cacheKeyMicroservicesData = "audi_lambda_microservices"
)

func (r *repository) GetAllMicroservicesName(ctx context.Context) (collections []string, err error) {
	collections, err = r.mongo.ListCollectionNames(ctx, bson.D{})
	if err != nil {
		r.logger.Errorf("error get all collections: %+v", err)
	}

	return
}

func (r *repository) GetMicroserviceDataByID(ctx context.Context, microserviceID string) (data *dto.PublicMicroserviceData, err error) {
	res, err := r.redis.Get(ctx, cacheKeyMicroservicesData).Result()
	if err != nil {
		r.logger.Errorf("error get microservice data for: %s, error: %+v", microserviceID, err)
		return
	}

	var microserviceData []dto.PublicMicroserviceData
	err = jsonutil.ConvertFromJSONSting(res, &microserviceData)
	if err != nil {
		r.logger.Errorf("error parse microservice data for: %s, error: %+v", microserviceID, err)
		return
	}

	for _, v := range microserviceData {
		if v.ID == microserviceID {
			data = &v
			return
		}
	}
	err = errors.ErrMicroserviceNotFound
	return
}

func (r *repository) GetMicroserviceAllEventDataCount(ctx context.Context, microserviceID string) (count dto.PublicMicroserviceDataCount, err error) {
	microserviceColls, err := r.mongo.ListCollectionNames(ctx, bson.D{{
		Key: "name",
		Value: bson.D{{
			Key: "$regex",
			Value: primitive.Regex{
				Pattern: fmt.Sprintf("^%s.*", microserviceID),
			},
		}},
	}})

	if err != nil {
		r.logger.Errorf("error list collections for: %s: %+v", microserviceID, err)
	}

	var total int64

	for _, collName := range microserviceColls {
		logLevel := microserviceutil.ExtractLogLevelFromID(collName)
		total, err = r.mongo.Collection(collName).CountDocuments(ctx, bson.D{})
		if err != nil {
			r.logger.Error("error getting count data from db for: %s, error: %+v", collName, err)
			return
		}

		count.TotalEventData += total
		countToProperAttribute(logLevel, total, &count)
	}
	return
}

func (r *repository) StoreMicroservicesData(ctx context.Context, microservicesData []dto.PublicMicroserviceData) (err error) {
	microservicesStr, err := jsonutil.ConvertToJSONString(microservicesData)
	if err != nil {
		r.logger.Errorf("error converting to json string: %+v", err)
		return
	}

	_, err = r.redis.Set(ctx, cacheKeyMicroservicesData, microservicesStr, 0).Result()
	if err != nil {
		r.logger.Errorf("error setting to cache for microservice data: %+v, error: %+v", microservicesData, err)
	}
	return
}

func (r *repository) FindMicroserviceEventData(ctx context.Context, microserviceID string, criteria indto.SearchEventCriteria) (events []dto.PublicEventData, err error) {
	events = make([]dto.PublicEventData, 0)

	microserviceColls, err := r.mongo.ListCollectionNames(ctx, bson.D{{
		Key: "name",
		Value: bson.D{{
			Key: "$regex",
			Value: primitive.Regex{
				Pattern: fmt.Sprintf("^%s.*", microserviceID),
			},
		}},
	}})

	if err != nil {
		r.logger.Errorf("error list collections for: %s: %+v", microserviceID, err)
	}

	for _, collName := range microserviceColls {
		filter := bson.D{}
		var curs *mongo.Cursor
		var currentEvents []model.EventLog

		if criteria.Level != nil {
			level := microserviceutil.ExtractLogLevelFromID(collName)
			if *criteria.Level != level {
				continue
			}
		}

		if criteria.Message != nil {
			msgCriteria := bson.D{{
				Key: "data",
				Value: bson.D{{
					Key: "$regex",
					Value: primitive.Regex{
						Pattern: fmt.Sprintf("^.*%s.*", *criteria.Message),
						Options: "i",
					},
				}},
			}}
			filter = append(filter, msgCriteria...)
		}

		if criteria.TimeStart != nil {
			timeCriteria := bson.D{{
				Key: "timestamp",
				Value: bson.D{{
					Key:   "$gte",
					Value: criteria.TimeStart,
				}},
			}}

			filter = append(filter, timeCriteria...)
		}

		if criteria.TimeEnd != nil {
			timeCriteria := bson.D{{
				Key: "timestamp",
				Value: bson.D{{
					Key:   "$lt",
					Value: criteria.TimeEnd,
				}},
			}}

			filter = append(filter, timeCriteria...)
		}

		curs, err = r.mongo.Collection(collName).Find(ctx, filter)
		if err != nil {
			r.logger.Errorf("error find mongodb for microserviceID: %s, search criteria: %+v, err: %+v", microserviceID, criteria, err)
			return
		}

		err = curs.All(ctx, &currentEvents)
		if err != nil {
			r.logger.Errorf("error decoding event for microserviceID: %s, search criteria: %+v, err: %+v", microserviceID, criteria, err)
			return
		}

		resp := microserviceutil.EventModelToRespDto(&currentEvents)
		if resp != nil {
			events = append(events, *resp...)
		}
	}
	return
}

func countToProperAttribute(level string, count int64, attr *dto.PublicMicroserviceDataCount) {
	switch level {
	case "info":
		attr.TotalInfoEventData += count
		break
	case "warn":
		attr.TotalWarnEventData += count
		break
	case "error":
		attr.TotalErrorEventData += count
	case "unknown":
		attr.TotalUnknownEventData += count
	}
}
