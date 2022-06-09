package microserviceutil

import "strings"

func ExtractLogLevelFromID(id string) (level string) {
	splittedID := strings.Split(id, "_")
	level = splittedID[len(splittedID)-1]
	return
}
