package wrapper

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Data struct{}

// ClientCredentials structure of client object
type ClientCredentials struct {
	Gateway, IPAddress, UserID string
}

// NewClient creation of client object
func NewClient(gateway, ipAddress, userID string) ClientCredentials {
	return ClientCredentials{
		Gateway:   gateway,
		IPAddress: ipAddress,
		UserID:    userID,
	}
}

// GetUser GET method to GET user information
func GetUser(c ClientCredentials) map[string]interface{} {
	_url := "http://uat-api.synapsefi.com/v3.1/users"

	client := &http.Client{}

	req, err := http.NewRequest("GET", _url, nil)

	req.Header.Set("x-sp-gateway", c.Gateway)
	req.Header.Set("x-sp-user-ip", c.IPAddress)
	req.Header.Set("x-sp-user", c.UserID)
	req.Header.Set("content-type", "application/json;charset=UTF-8")

	if err != nil {
		fmt.Println(err)
	}

	resp, err := client.Do(req)

	if err != nil {
		fmt.Println(err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		fmt.Println(err)
	}

	var data interface{}
	json.Unmarshal(body, &data)

	jsonData, ok := data.(map[string]interface{})

	if ok != false {
		jsonData["id"] = c.UserID
	}

	return jsonData
}
