package main

import (
	"encoding/json"
)

// Config 结构体定义
type Config struct {
	Mode   string      `json:"mode"`
	Config interface{} `json:"config"`
}

// BreathConfig 结构体定义
type BreathConfig struct {
	Onboard    int     `json:"onboard"`
	Brightness float64 `json:"brightness"`
	Speed      float64 `json:"speed"`
	Color      []int   `json:"color"`
}

// 生成 BreathConfig 对应的 JSON 格式字符串的方法
func (b BreathConfig) toJSONString() (string, error) {
	config := Config{
		Mode:   "breath",
		Config: b,
	}

	jsonBytes, err := json.Marshal(config)
	if err != nil {
		return "", err
	}
	return string(jsonBytes), nil
}

// GradientConfig 结构体定义
type GradientConfig struct {
	Onboard    int     `json:"onboard"`
	Brightness float64 `json:"brightness"`
	Speed      float64 `json:"speed"`
	ColorFrom  []int   `json:"color_from"`
	ColorTo    []int   `json:"color_to"`
}

// 生成 GradientConfig 对应的 JSON 格式字符串的方法
func (g GradientConfig) toJSONString() (string, error) {
	config := Config{
		Mode:   "gradient",
		Config: g,
	}

	jsonBytes, err := json.Marshal(config)
	if err != nil {
		return "", err
	}
	return string(jsonBytes), nil
}
