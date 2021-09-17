package go2struct

import (
	"testing"
)

func TestJSON2Struct(t *testing.T) {
	temp := `
	{
	 "code": 200,
	 "success": true,
	 "msg": {
		"content": "Hello"
	 }
	}
	`

	res, err := JSON2Struct([]byte(temp))
	if err != nil {
		t.Fatalf("json to struct failed, error: \n%v", err)
		return
	}
	t.Logf("json to struct success, result: \n%s", res)
}
