package main

import (
	"encoding/json"
)

func ParseUdp(jsonString string) (string, error) {
	// 定义一个结构体来解析 JSON
	var jsonData map[string]json.RawMessage
	if err := json.Unmarshal([]byte(jsonString), &jsonData); err != nil {
		// fmt.Println("JSON 解析错误:", err)
		return "", err
	}

	var dataType string
	if err := json.Unmarshal(jsonData["type"], &dataType); err != nil {
		// fmt.Println("类型解析错误:", err)
		return "", err
	}

	dataContent := string(jsonData["data"])

	return dataContent, nil
}
