package main

import (
	"fmt"
	"net"
)

func Udp_send(targetIP, targetPort, data string) {

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

func HandleUDP() {
	// 监听UDP端口
	addr, err := net.ResolveUDPAddr("udp", ":6688")
	if err != nil {
		fmt.Println("Error resolving UDP address:", err)
		return
	}

	conn, err := net.ListenUDP("udp", addr)
	if err != nil {
		fmt.Println("Error listening to UDP:", err)
		return
	}
	defer conn.Close()

	fmt.Println("Listening for UDP on", addr)

	buffer := make([]byte, 1024)

	for {
		// 读取UDP数据包
		n, addr, err := conn.ReadFromUDP(buffer)
		if err != nil {
			fmt.Println("Error reading UDP:", err)
			continue
		}

		fmt.Printf("Received UDP message from %s: %s\n", addr.String(), string(buffer[:n]))
	}
}
