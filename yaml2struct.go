package go2struct

import "gopkg.in/yaml.v2"

// YAML2Struct convert yaml to struct
func YAML2Struct(name string, data []byte, args ...string) ([]byte, error) {
	var m map[string]interface{}

	if err := yaml.Unmarshal(data, &m); err != nil {
		return nil, err
	}

	return Map2Struct(name, m, args...), nil
}
