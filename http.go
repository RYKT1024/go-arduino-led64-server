package main

import (
	"fmt"
	"net/http"
)

func ConfigHandler(w http.ResponseWriter, r *http.Request) {
	// 使用通道来获取 HandleUDP 函数的返回值
	resultChannel := make(chan string)

	// 启动 HandleUDP 函数的 goroutine
	go func() {
		result := HandleUDP()
		// 将结果发送到通道
		resultChannel <- result
	}()

	UdpSend("192.168.31.241", "1122", `{"mode":"get_onboard"}`)

	// 从通道中获取返回值
	result := <-resultChannel

	// 打印返回值
	fmt.Println("HandleUDP result:", result)
	fmt.Fprint(w, result)
}

//
