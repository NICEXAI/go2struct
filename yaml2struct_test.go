package go2struct

import "testing"

func TestYAML2Struct(t *testing.T) {
	temp := `
code: 200
success: true
msg:
 content: Hello
`

	res, err := YAML2Struct([]byte(temp))
	if err != nil {
		t.Fatalf("yaml to struct failed, error: \n%v", err)
		return
	}
	t.Logf("yaml to struct success, result: \n%s", res)
}