package helpers

import "encoding/json"

/*
 * Load file JSON ke storage
 *
 */
func LoadFromJSON(content []byte, v interface{}) error {
	return json.Unmarshal(content, v)
}

/*
 * Menyimpan data ke file JSON
 *
 */
func SaveToJSON(v interface{}) ([]byte, error) {
	return json.Marshal(v)
}