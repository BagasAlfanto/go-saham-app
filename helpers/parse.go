package helpers

import "encoding/json"

func LoadFromJSON(content []byte, v interface{}) error {
	return json.Unmarshal(content, v)
}
