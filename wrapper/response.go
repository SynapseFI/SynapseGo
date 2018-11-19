package wrapper

import (
	"encoding/json"
	"fmt"
)

func read(data []byte) map[string]interface{} {
	d := make(map[string]interface{})
	err := json.Unmarshal(data, &d)

	if err != nil {
		errorLog(err)
	}

	return d
}

func multiData(data []byte, setting string) map[string]interface{} {
	body := make(map[string]interface{})
	d := read(data)

	body["limit"] = d["limit"]
	body["page"] = d["page"]
	body["pageCount"] = d["page_count"]

	switch setting {
	case "nodes":
		body["nodeCount"] = d["node_count"]
		body["nodesList"] = list(d["nodes"].(map[string]interface{}), "node")
	case "subscriptions":
		body["subscriptionsCount"] = d["subscriptions_count"]
		body["subsList"] = list(d["subscriptions"], "subscription")
	case "transactions":
		body["transCount"] = d["trans_count"]
		body["transList"] = list(d["trans"], "transaction")
	case "users":
		fmt.Println(d)
		body["usersCount"] = d["users_count"]
		body["usersList"] = list(d["users"], "user")
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
			v["id"] = d[i].(map[string]interface{})["_id"]
			v["url"] = d[i]
			v["payload"] = d[i]

			list = append(list, v)
		}

	case "user":
		for i := 0; i < len(d); i++ {
			v := make(map[string]interface{})
			v["id"] = d[i].(map[string]interface{})["_id"]
			v["fullDehydrate"] = true
			v["payload"] = d[i]

			list = append(list, v)
		}
	default:
	}

	return list
}

func singleData(value []byte, setting string) map[string]interface{} {
	body := make(map[string]interface{})
	v := read(value)

	switch setting {
	case "node":

	case "subscription":

	case "transaction":

	case "user":
		body["id"] = v["id"]
		body["fullDehydrate"] = true
		body["payload"] = v
	}

	return body
}
