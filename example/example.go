package main

import (
	"encoding/json"
	"fmt"
	"github.com/NICEXAI/go2struct"
)

type Message struct {
	Code    int  `json:"code"`
	Success bool `json:"success"`
	Msg     struct {
		Content string
	} `json:"msg"`
}

func main() {
	var options map[string]interface{}

	message := Message{
		Code: 200,
		Msg:  struct{ Content string }{Content: "Hello"},
	}

	bData, err := json.Marshal(message)
	if err != nil {
		fmt.Printf("json marshal failed, error: %v", err)
		return
	}
	_ = json.Unmarshal(bData, &options)

	//map to struct
	structTxt := go2struct.Map2Struct(options)

	fmt.Printf("map to struct success, result: \n%s", structTxt)

	// json to struct
	jsonTemp := `
	{
	 "code": 200,
	 "success": true,
	 "msg": {
		"content": "Hello"
	 }
	}
	`

	res, err := go2struct.JSON2Struct([]byte(jsonTemp))
	if err != nil {
		fmt.Printf("json to struct failed, error: \n%v", err)
		return
	}
	fmt.Printf("json to struct success, result: \n%s", res)

	// yaml to struct
	temp := `
code: 200
success: true
msg:
 content: Hello
`

	res, err = go2struct.YAML2Struct([]byte(temp))
	if err != nil {
		fmt.Printf("yaml to struct failed, error: \n%v", err)
		return
	}
	fmt.Printf("yaml to struct success, result: \n%s", res)
}
