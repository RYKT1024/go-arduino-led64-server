package main

import (
	"fmt"
	"io"
	"net/http"
)

var IP string = "192.168.31.241"
var PORT string = "1122"

func ConfigHandler(w http.ResponseWriter, r *http.Request) {
	// 使用通道来获取 HandleUDP 函数的返回值
	resultChannel := make(chan string)

	// 启动 HandleUDP 函数的 goroutine
	go func() {
		result := HandleUDP()
		// 将结果发送到通道
		resultChannel <- result
	}()

	UdpSend(IP, PORT, `{"mode":"get_onboard"}`)

	// 从通道中获取返回值
	result := <-resultChannel

	// 打印返回值
	fmt.Println("UDP config get successfully.")
	fmt.Fprint(w, result)
}

func SetHandler(w http.ResponseWriter, r *http.Request) {

	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading request body", http.StatusInternalServerError)
		return
	}

	bodyString := string(body)
	UdpSend(IP, PORT, bodyString)
	fmt.Fprint(w, bodyString)
}

func StatusHandler(w http.ResponseWriter, r *http.Request) {
	resultChannel := make(chan string)

	go func() {
		result := HandleUDP()
		resultChannel <- result
	}()

	UdpSend(IP, PORT, `{"mode":"status"}`)

	// 从通道中获取返回值
	result := <-resultChannel

	// 打印返回值
	fmt.Println("UDP status get successfully.")
	fmt.Fprint(w, result)
}
