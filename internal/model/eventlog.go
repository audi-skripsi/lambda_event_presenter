package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type EventLog struct {
	ObjectID  primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	UID       string             `json:"uid" bson:"uid"`
	Level     string             `json:"level" bson:"level"`
	AppName   string             `json:"app_name" bson:"app_name"`
	Data      interface{}        `json:"data" bson:"data"`
	Timestamp time.Time          `json:"timestamp" bson:"timestamp"`
}
