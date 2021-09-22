package go2struct

import (
	"fmt"
	"github.com/NICEXAI/go2struct/util"
	"strings"
)

const (
	structFirst    = "type "
	structLast     = "}"
	structStartTag = "%s struct {\n"
	structEndTag   = "%s} `json:\"%s\"`\n"
	structSpace    = "    "
	structFieldTag = "%s%s %s `json:\"%s\"`\n"
)

// Map2Struct convert map to struct
func Map2Struct(name string, m map[string]interface{}) []byte {
	cellNodes := convertMapToCellNode(name, m, 0)
	return []byte(strings.Join(cellNodes, ""))
}

func convertMapToCellNode(name string, m map[string]interface{}, tier int) (cn []string) {
	if tier == 0 {
		cn = append(cn, structFirst)
	}
	wrapperSpace := getSpaceByTier(tier - 1)
	cn = append(cn, fmt.Sprintf(structStartTag, wrapperSpace+util.UnderscoreToUpperCamelCase(name)))

	for field, val := range m {
		fName := util.UnderscoreToUpperCamelCase(field)
		fType := ""

		switch val.(type) {
		case float64:
			if strings.Contains(fmt.Sprintf("%v", val), ".") {
				fType = "float64"
			} else {
				fType = "int"
			}
		case int:
			fType = "int"
		case bool:
			fType = "bool"
		case string:
			fType = "string"
		case map[string]interface{}:
			fType = "struct"
		case map[interface{}]interface{}:
			//convert map[interface{}]interface{} to map[string]interface{}
			newVal := make(map[string]interface{})
			for k, v := range val.(map[interface{}]interface{}) {
				strKey := fmt.Sprintf("%v", k)
				newVal[strKey] = v
			}
			val = newVal
			fType = "struct"
		}

		if fType != "struct" {
			cn = append(cn, fmt.Sprintf(structFieldTag, getSpaceByTier(tier), fName, fType, util.UpperCamelCaseToUnderscore(field)))
		} else {
			child := convertMapToCellNode(field, val.(map[string]interface{}), tier+1)
			cn = append(cn, child...)
		}
	}

	if tier == 0 {
		cn = append(cn, structLast)
	} else {
		cn = append(cn, fmt.Sprintf(structEndTag, wrapperSpace, name))
	}
	return cn
}

func getSpaceByTier(tier int) (s string) {
	for i := 0; i < tier+1; i++ {
		s += structSpace
	}
	return s
}
