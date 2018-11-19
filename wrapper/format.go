package wrapper

import (
	"encoding/json"
)

func response(data []byte, setting string) map[string]interface{} {
	body := make(map[string]interface{})
	d := make(map[string]interface{})
	err := json.Unmarshal(data, &d)

	if err != nil {
		errorLog(err)
	}

	body["limit"] = d["limit"]
	body["page"] = d["page"]
	body["pageCount"] = d["page_count"]

	switch setting {
	case "nodes":
		body["nodeCount"] = d["node_count"]
		body["nodesList"] = list(d["nodes"].(map[string]interface{}), "node")
	case "subscriptins":
		body["subscriptionsCount"] = d["subscriptions_count"]
		body["subsList"] = list(d["subscriptions"], "subscription")
	case "transactions":
		body["transCount"] = d["trans_count"]
		body["transList"] = d
	case "users":
		body["usersCount"] = d["users_count"]
		body["usersList"] = d
	}

	return body
}

func list(data interface{}, setting string) []interface{} {
	var list []interface{}
	d := data.([]interface{})

	switch setting {
	case "subscription":
		for i := 0; i < len(d); i++ {
			v := make(map[string]interface{})
			v["id"] = d[i]
			v["url"] = d[i]
			v["payload"] = d[i]

			list = append(list, v)
		}
		// r["id"] = d["_id"]
		// r["payload"] = d
		// r["url"] = d["url"]

	default:
	}

	return list
}
