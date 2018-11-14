package main

import (
	"api-wrapper/wrapper"
	"fmt"
)

func main() {
	// DEVELOPER WILL CREATE THE FOLLOWING:
	client := wrapper.NewClient(
		"client_id_pSdgBAaNfvRCcFmqK682tIhXOyD5iEV0rJx3nskQ|client_secret_7GfSOVxD5seYHRC8o2X0grlvZtd9azmhMpn3U1Ju",
		"127.0.0.1",
		"|5beb505292571b00a14dd31a",
	)

	fmt.Println(wrapper.GetUser(client))
	// wrapper.GetUser(client)

}
