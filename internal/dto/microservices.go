package dto

import "time"

type SearchEventCriteria struct {
	Level     *string
	Message   *string
	TimeStart *time.Time
	TimeEnd   *time.Time
}
