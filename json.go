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

// // 创建 BreathConfig 结构体实例
// breathConfig := BreathConfig{
// 	Onboard:    0,
// 	Brightness: 0.2,
// 	Speed:      0.3,
// 	Color:      []int{255, 0, 120},
// }

// // 使用 toJSONString 方法将 BreathConfig 转换为 JSON 字符串
// breathJSON, err := breathConfig.toJSONString()
// if err != nil {
// 	fmt.Println("Error converting BreathConfig to JSON:", err)
// 	return
// }

// // 创建 GradientConfig 结构体实例
// gradientConfig := GradientConfig{
// 	Onboard:    0,
// 	Brightness: 0.4,
// 	Speed:      0.8,
// 	ColorFrom:  []int{255, 0, 120},
// 	ColorTo:    []int{0, 255, 0},
// }

// // 使用 toJSONString 方法将 GradientConfig 转换为 JSON 字符串
// gradientJSON, err := gradientConfig.toJSONString()
// if err != nil {
// 	fmt.Println("Error converting GradientConfig to JSON:", err)
// 	return
// }
