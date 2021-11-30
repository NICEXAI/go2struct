package go2struct

import (
	"fmt"
	"github.com/NICEXAI/go2struct/util"
	"strings"
)

const (
	structFirst    = "type "
	structLast     = "}"
	structStartTag = "%s %sstruct {\n"
	structEndTag   = "%s} `json:\"%s\"`\n"
	structSpace    = "    "
	structFieldTag = "%s%s %s `json:\"%s\"`\n"
)

// Map2Struct convert map to struct
func Map2Struct(name string, m map[string]interface{}) []byte {
	cellNodes := convertMapToCellNode(name, m, false, 0)
	return []byte(strings.Join(cellNodes, ""))
}

func convertMapToCellNode(name string, m map[string]interface{}, isSlice bool, tier int) (cn []string) {
	if tier == 0 {
		cn = append(cn, structFirst)
	}
	wrapperSpace := getSpaceByTier(tier - 1)
	if isSlice {
		cn = append(cn, fmt.Sprintf(structStartTag, wrapperSpace+util.UnderscoreToUpperCamelCase(name), "[]"))
	} else {
		cn = append(cn, fmt.Sprintf(structStartTag, wrapperSpace+util.UnderscoreToUpperCamelCase(name), ""))
	}

	for field, val := range m {
		fName := util.UnderscoreToUpperCamelCase(field)
		fType := getFiledType(val)

		if fType == "struct" {
			if list, ok := val.(map[interface{}]interface{}); ok {
				//convert map[interface{}]interface{} to map[string]interface{}
				newVal := make(map[string]interface{})
				for k, v := range list {
					strKey := fmt.Sprintf("%v", k)
					newVal[strKey] = v
				}
				val = newVal
			}
		}

		if fType != "struct" && fType != "slice" {
			cn = append(cn, fmt.Sprintf(structFieldTag, getSpaceByTier(tier), fName, fType, util.UpperCamelCaseToUnderscore(field)))
		}

		if fType == "struct" {
			child := convertMapToCellNode(field, val.(map[string]interface{}), false, tier+1)
			cn = append(cn, child...)
		}

		if fType == "slice" {
			subList, _ := val.([]interface{})
			if len(subList) > 0 {
				fSubType := getFiledType(subList[0])
				if fSubType != "struct" && fSubType != "slice" {
					cn = append(cn, fmt.Sprintf(structFieldTag, getSpaceByTier(tier), fName, "[]"+fSubType, util.UpperCamelCaseToUnderscore(field)))
				}
				if fSubType == "struct" {
					child := convertMapToCellNode(field, subList[0].(map[string]interface{}), true, tier+1)
					cn = append(cn, child...)
				}
			}
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

func getFiledType(filed interface{}) string {
	switch filed.(type) {
	case float64:
		if strings.Contains(fmt.Sprintf("%v", filed), ".") {
			return "float64"
		}
		return "int"
	case int:
		return "int"
	case bool:
		return "bool"
	case string:
		return "string"
	case map[string]interface{}:
		return "struct"
	case map[interface{}]interface{}:
		return "struct"
	case []interface{}:
		return "slice"
	default:
		return ""
	}
}
