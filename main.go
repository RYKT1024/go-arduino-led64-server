package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func main() {

	// 创建mux路由器
	router := mux.NewRouter()

	// 添加HTTP路由
	router.HandleFunc("/config", ConfigHandler).Methods("GET")
	router.HandleFunc("/set", SetHandler).Methods("POST")
	router.HandleFunc("/set_IP", SetIP_Handler).Methods("POST")

	// 启动HTTP服务器
	fmt.Println("Listening for HTTP on :6789")

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "OPTIONS"},
		AllowedHeaders:   []string{"Content-Type", "Origin"},
		AllowCredentials: true,
	})

	handler := c.Handler(router)
	http.ListenAndServe(":6789", handler)
}
