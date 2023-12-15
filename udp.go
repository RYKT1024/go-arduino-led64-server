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
