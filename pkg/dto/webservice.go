package dto

type PublicMicroserviceData struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type PublicMicroserviceDataCount struct {
	TotalEventData        int64 `json:"totalEventData"`
	TotalInfoEventData    int64 `json:"totalInfoEventData"`
	TotalWarnEventData    int64 `json:"totalWarnEventData"`
	TotalErrorEventData   int64 `json:"totalErrorEventData"`
	TotalUnknownEventData int64 `json:"totalUnknownEventData"`
}

type PublicEventData struct {
	UID         string `json:"uid"`
	Level       string `json:"level"`
	Message     string `json:"message"`
	ServiceName string `json:"serviceName"`
	Timestamp   int64  `json:"timestamp"`
}

type PublicMicroservicesNameResponse struct {
	Microservices []PublicMicroserviceData `json:"microservices"`
}

type PublicMicroserviceAnalyticsResponse struct {
	EventDataCount PublicMicroserviceDataCount `json:"eventDataCount"`
}

type PublicMicroserviceEventSearchResponse struct {
	TotalEventData int64             `json:"totalEventData"`
	Events         []PublicEventData `json:"events"`
}
