package wrapper

import (
	"encoding/json"
)

func format(data []byte) map[string]interface{} {
	m := make(map[string]interface{})
	err := json.Unmarshal(data, &m)

	if err != nil {
		errorLog(err)
	}

	return m
}
