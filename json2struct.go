package go2struct

import "encoding/json"

// JSON2Struct convert json to struct
func JSON2Struct(name string, data []byte) ([]byte, error) {
	var m map[string]interface{}

	if err := json.Unmarshal(data, &m); err != nil {
		return nil, err
	}

	return Map2Struct(name, m), nil
}
