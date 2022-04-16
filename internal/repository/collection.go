package repository

import (
	"context"

	"github.com/audi-skripsi/lambda_event_presenter/internal/constant"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var indexesModels = []mongo.IndexModel{
	{
		Keys: bson.D{
			{
				Key:   "uid",
				Value: 1,
			},
		},
	},
	{
		Keys: bson.D{
			{
				Key:   "level",
				Value: 1,
			},
		},
	},
	{
		Keys: bson.D{
			{
				Key:   "level",
				Value: 1,
			},
			{
				Key:   "timestamp",
				Value: 1,
			},
		},
	},
	{
		Keys: bson.D{
			{
				Key:   "timestamp",
				Value: 1,
			},
		},
	},
}

func (r *repository) SegragateCollection(name string) (err error) {
	opt1 := r.redis.HGet(
		context.Background(),
		constant.RedisKeyCollectionCheckup,
		name,
	)

	res, err := opt1.Result()
	if err != nil {
		r.logger.Errorf("error interacting with redis for %s: %+v", name, err)
	} else {
		if res == "1" {
			r.logger.Infof("collection found in cache")
			return
		}
	}

	cmd := r.redis.HMSet(context.Background(), constant.RedisKeyCollectionCheckup, name, "1")
	if cmd.Err() != nil {
		r.logger.Errorf("error hset redis for %s: %+v", name, cmd.Err())
	}

	err = nil

	indexes, err := r.mongo.Collection(name).Indexes().CreateMany(
		context.Background(),
		indexesModels,
	)
	if err != nil {
		r.logger.Errorf("error create many indexes: %+v", err)
		return
	}

	r.logger.Infof("indexes created: %+v", indexes)

	opt2 := r.redis.HMSet(context.Background(), constant.RedisKeyCollectionCheckup, name, "1")
	if opt2.Err() != nil {
		r.logger.Errorf("error setting new redis: %+v", err)
	}

	return
}
