package go2struct

import "gopkg.in/yaml.v2"

// YAML2Struct convert yaml to struct
func YAML2Struct(data []byte) (string, error) {
	var m map[string]interface{}

	if err := yaml.Unmarshal(data, &m); err != nil {
		return "", err
	}

	return Map2Struct(m), nil
}
