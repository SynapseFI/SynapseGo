package wrapper

import (
	"encoding/json"
)

/********** METHODS **********/

func data(data []byte) map[string]interface{} {
	body := make(map[string]interface{})
	d := read(data)

	body["payload"] = d

	return body
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

		if d["nodes"] != nil {
			body["nodesList"] = list(d["nodes"].(map[string]interface{}), "node")
		} else {
			body["nodesList"] = nil
		}

	case "subscriptions":
		body["subscriptionsCount"] = d["subscriptions_count"]
		body["subsList"] = list(d["subscriptions"], "subscription")
	case "transactions":
		body["transCount"] = d["trans_count"]
		body["transList"] = list(d["trans"], "transaction")
	case "users":
		body["usersCount"] = d["users_count"]
		body["usersList"] = list(d["users"], "user")
	}

	return body
}

func singleData(value map[string]interface{}, setting string) map[string]interface{} {
	body := make(map[string]interface{})

	switch setting {
	case "node":
		body["id"] = value["_id"]
		body["userID"] = value["user_id"]
		body["fullDehydrate"] = true
		body["payload"] = value

	case "subscription":
		body["id"] = value["_id"]
		body["url"] = value["url"]
		body["payload"] = value["payload"]

	case "transaction":
		body["id"] = value["_id"]
		body["payload"] = value

	case "user":
		body["id"] = value["id"]
		body["fullDehydrate"] = true
		body["payload"] = value
	}

	return body
}

/********** HELPERS **********/

func list(data interface{}, setting string) []interface{} {
	var list []interface{}

	if data != nil {
		d := data.([]interface{})

		for i := 0; i < len(d); i++ {
			k := d[i].(map[string]interface{})
			v := singleData(k, setting)

			list = append(list, v)
		}
	}

	return list
}

func read(data []byte) map[string]interface{} {
	d := make(map[string]interface{})
	err := json.Unmarshal(data, &d)

	if err != nil {
		errorLog(err)
	}

	return d
}
