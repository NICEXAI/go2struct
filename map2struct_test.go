package go2struct

import (
	"encoding/json"
	"testing"
)

type Message struct {
	Code    int  `json:"code"`
	Success bool `json:"success"`
	Msg     struct {
		Content string
	} `json:"msg"`
}

func TestMap2Struct(t *testing.T) {
	var options map[string]interface{}

	message := Message{
		Code: 200,
		Msg:  struct{ Content string }{Content: "Hello"},
	}

	bData, err := json.Marshal(message)
	if err != nil {
		t.Fatalf("json marshal failed, error: %v", err)
		return
	}
	_ = json.Unmarshal(bData, &options)
	t.Logf("map to struct success, result: \n%s", Map2Struct(options))
}
