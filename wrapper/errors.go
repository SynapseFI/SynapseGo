package wrapper

import (
	"fmt"
)

func handleError(err map[string]interface{}) {
	fmt.Println(err)
	panic(err["error"].(map[string]interface{})["en"].(string))
}
