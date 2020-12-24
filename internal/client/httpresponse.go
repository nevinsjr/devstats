package client

import (
	"encoding/json"
	"net/http"
)

type ResponsePostProcessor func(response *http.Response) error

func Deserializer(model interface{}) ResponsePostProcessor {
	return func(response *http.Response) error {
		return json.NewDecoder(response.Body).Decode(model)
	}
}
