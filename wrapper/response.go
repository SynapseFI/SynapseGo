package wrapper

import (
	"encoding/json"
	"strconv"
)

/********** METHODS **********/

func response(data []byte, key ...string) map[string]interface{} {
	return read(data)
}

func responseMulti(data []byte, setting string) map[string]interface{} {
	body := make(map[string]interface{})
	d := read(data)

	body["limit"] = d["limit"]
	body["page"] = d["page"]
	body["pageCount"] = d["page_count"]

	switch setting {
	case "nodes":
		body["nodeCount"] = d["node_count"]
		body["nodesList"] = list(d["nodes"], "node")

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

func responseSingle(data []byte, setting string) map[string]interface{} {
	body := make(map[string]interface{})
	d := read(data)

	switch setting {
	case "node":
		body["id"] = d["_id"]
		body["userID"] = d["user_id"]
		body["fullDehydrate"] = "no"
		body["payload"] = d

	case "subscription":
		body["id"] = d["_id"]
		body["url"] = d["url"]
		body["payload"] = d

	case "transaction":
		body["id"] = d["_id"]
		body["payload"] = d

	case "user":
		body["id"] = d["_id"]
		body["payload"] = d
	}

	return body
}

/********** HELPERS **********/

func list(data interface{}, setting string) []interface{} {
	var list []interface{}

	if data != nil {
		d := data.([]interface{})

		for i := 0; i < len(d); i++ {
			k, err := json.Marshal(d[i].(map[string]interface{}))

			if err != nil {
				errorLog(err)
			}

			v := responseSingle(k, setting)
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

func queryString(params []map[string]interface{}) string {
	var query string

	if len(params) > 0 {
		for k := range params[0] {
			if k == "query" || k == "show_refresh_tokens" {
				query += k + "=" + params[0][k].(string) + "&"
			} else if k == "page" || k == "per_page" {
				query += k + "=" + strconv.Itoa(params[0][k].(int)) + "&"
			}
		}
	}

	return query
}
