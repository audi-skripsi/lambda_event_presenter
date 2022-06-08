package microserviceutil

import (
	"strings"

	"github.com/audi-skripsi/lambda_event_presenter/pkg/dto"
)

// ParseMicroserviceIDFromCollection parses microservice id from collection
// that is already formatted in a such to fit bucketing pattern.
func ParseMicroserviceIDFromCollection(coll string) (id string) {
	separatedString := strings.Split(coll, "_")
	if len(separatedString) < 4 {
		return coll
	}
	microserviceID := separatedString[0:(len(separatedString) - 3)]
	return strings.Join(microserviceID, "_")
}

// ParseMicroserviceIDFromCollection parses microservice name from collection
// that is already formatted in a such to fit bucketing pattern.
func ParseMicroserviceNameFromCollection(coll string) (name string) {
	separatedString := strings.Split(coll, "_")
	if len(separatedString) < 4 {
		return coll
	}
	microserviceID := separatedString[0:(len(separatedString) - 3)]
	return strings.Join(microserviceID, " ")
}

// ParseDataFromCollection parses microservice data from collection to public dto
func ParseDataFromCollections(collsName []string) (microserviceData []dto.PublicMicroserviceData) {
	uniqueMicroservices := make(map[string]dto.PublicMicroserviceData)

	for _, v := range collsName {
		id := ParseMicroserviceIDFromCollection(v)
		name := ParseMicroserviceNameFromCollection(v)
		if _, ok := uniqueMicroservices[id]; !ok {
			uniqueMicroservices[id] = dto.PublicMicroserviceData{
				ID:   id,
				Name: name,
			}
		}
	}

	for _, v := range uniqueMicroservices {
		microserviceData = append(microserviceData, v)
	}

	return
}
