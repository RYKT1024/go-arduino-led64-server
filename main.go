package main

import "fmt"

func main() {
	//Udp_send("192.168.31.241", "1122", `{"mode":"gradient", "config":{"onboard":0, "brightness":0.1, "speed":0.8, "color_from":[255, 0, 0], "color_to":[0, 255, 0]}}`)

	// 创建 BreathConfig 结构体实例
	breathConfig := BreathConfig{
		Onboard:    0,
		Brightness: 0.2,
		Speed:      0.3,
		Color:      []int{255, 0, 120},
	}

	// 使用 toJSONString 方法将 BreathConfig 转换为 JSON 字符串
	breathJSON, err := breathConfig.toJSONString()
	if err != nil {
		fmt.Println("Error converting BreathConfig to JSON:", err)
		return
	}

	// 打印生成的 BreathConfig JSON 字符串
	fmt.Println(breathJSON)

	// 创建 GradientConfig 结构体实例
	gradientConfig := GradientConfig{
		Onboard:    0,
		Brightness: 0.4,
		Speed:      0.8,
		ColorFrom:  []int{255, 0, 120},
		ColorTo:    []int{0, 255, 0},
	}

	// 使用 toJSONString 方法将 GradientConfig 转换为 JSON 字符串
	gradientJSON, err := gradientConfig.toJSONString()
	if err != nil {
		fmt.Println("Error converting GradientConfig to JSON:", err)
		return
	}

	// 打印生成的 GradientConfig JSON 字符串
	fmt.Println(gradientJSON)

	Udp_send("192.168.31.241", "1122", gradientJSON)
}
