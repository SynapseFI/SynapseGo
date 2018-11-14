package main

import (
	"api-wrapper/wrapper"
	"fmt"
)

func main() {
	// DEVELOPER WILL DO THE FOLLOWING:
	client := wrapper.NewClient(
		"client_id_pSdgBAaNfvRCcFmqK682tIhXOyD5iEV0rJx3nskQ|client_secret_7GfSOVxD5seYHRC8o2X0grlvZtd9azmhMpn3U1Ju",
		"127.0.0.1",
		"|5beb505292571b00a14dd31a",
	)

	// data, err := wrapper.GetUsers(client)
	// data, err := wrapper.GetUser(client, "5bec6ebebaabfc00ab168fa0")
	var newUserDataTest = []byte(`{
		"logins":       "test",
		"phoneNumbers": "2035551234",
		"legalNames":   "none",}`)
	data, err := wrapper.CreateUser(client, newUserDataTest)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(data))

}
