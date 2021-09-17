package go2struct

import "encoding/json"

// JSON2Struct convert json to struct
func JSON2Struct(data []byte) (string, error) {
	var m map[string]interface{}

	if err := json.Unmarshal(data, &m); err != nil {
		return "", err
	}

	return Map2Struct(m), nil
}
