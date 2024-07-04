package parser

import (
	"encoding/json"
)

// ParseDataToType parses data into the target type specified by targetType.
// data: The input data to be parsed, typically JSON data.
// targetType: A pointer to the target type into which the data will be unmarshaled.
// Returns the parsed target type and any error encountered during parsing.
func ParseDataToType(data interface{}, targetType interface{}) (interface{}, error) {
	jsonData, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(jsonData, &targetType)
	if err != nil {
		return nil, err
	}

	return targetType, nil
}
