package synapse

import "encoding/json"

func read(data []byte) map[string]interface{} {
	d := make(map[string]interface{})
	err := json.Unmarshal(data, &d)

	if err != nil {
		panic(err)
	}

	return d
}
