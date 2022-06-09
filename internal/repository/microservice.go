package repository

import (
	"context"
	"fmt"

	"github.com/audi-skripsi/lambda_event_presenter/internal/util/microserviceutil"
	"github.com/audi-skripsi/lambda_event_presenter/pkg/dto"
	"github.com/audi-skripsi/lambda_event_presenter/pkg/errors"
	"github.com/audi-skripsi/lambda_event_presenter/pkg/util/jsonutil"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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
