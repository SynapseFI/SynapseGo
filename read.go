package synapse

import (
	"encoding/json"
)

func readStream(data []byte) map[string]interface{} {
	d := make(map[string]interface{})
	err := json.Unmarshal(data, &d)

	// if data is an empty stream this will cause an unmarshal error
	if err != nil {
		panic(err)
	}

	return d
}
