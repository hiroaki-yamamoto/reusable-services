package errors

import (
	"encoding/json"
	"log"
	"strings"
)

// NotFound indicates the model wasn't found / rotted.
type NotFound struct {
	Metadata map[string]interface{}
}

func (me *NotFound) Error() string {
	var builder strings.Builder
	enc := json.NewEncoder(&builder)
	if err := enc.Encode(me.Metadata); err != nil {
		builder.Reset()
		log.Println("JSON Encoding Error", err)
	}
	return "Not Found: " + builder.String()
}
