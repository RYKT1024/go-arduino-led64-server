package main

import (
	"fmt"
	"net"
)

func UdpSend(targetIP, targetPort, data string) {

	// 构建目标地址
	targetAddr, err := net.ResolveUDPAddr("udp", targetIP+":"+targetPort)
	if err != nil {
		fmt.Println("Error resolving target address:", err)
	}

	// 创建UDP连接
	conn, err := net.DialUDP("udp", nil, targetAddr)
	if err != nil {
		fmt.Println("Error creating UDP connection:", err)
	}
	defer conn.Close()

	// 将字符串转换为字节数组，并发送UDP数据包
	byteData := []byte(data)
	_, err = conn.Write(byteData)
	if err != nil {
		fmt.Println("Error sending UDP packet:", err)
	}

	fmt.Println("UDP packet sent successfully.")
}

func HandleUDP() string {
	// 监听UDP端口
	addr, err := net.ResolveUDPAddr("udp", ":6688")
	if err != nil {
		fmt.Println("Error resolving UDP address:", err)
		return ""
	}

	conn, err := net.ListenUDP("udp", addr)
	if err != nil {
		fmt.Println("Error listening to UDP:", err)
		return ""
	}
	defer conn.Close()

	fmt.Println("Listening for UDP on", addr)

	buffer := make([]byte, 2048)
	var json_data string

	for {
		// 读取UDP数据包
		n, addr, err := conn.ReadFromUDP(buffer)
		if err != nil {
			fmt.Println("Error reading UDP:", err)
			continue
		}

		// 创建新的切片，只包含有效数据
		data := make([]byte, n)
		copy(data, buffer[:n])
		json_data = string(data)

		fmt.Printf("Received UDP message from %s: %s\n", addr.String(), json_data)
		dataType, dataContent, err := ParseUdp(json_data)
		if err != nil {
			fmt.Println("Error parsing UDP:", err)
			continue
		}

		if dataType == "config" {
			return dataContent
		}

	}
}
