package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	// 创建mux路由器
	router := mux.NewRouter()

	// 添加HTTP路由
	router.HandleFunc("/", HandleHTTP)

	// 启动HTTP服务器
	go func() {
		fmt.Println("Listening for HTTP on :6789")
		err := http.ListenAndServe(":6789", router)
		if err != nil {
			fmt.Println("Error starting HTTP server:", err)
		}
	}()

	// 启动UDP服务器
	go HandleUDP()

	// 阻塞主goroutine，使程序持续运行
	select {}
}
