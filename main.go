package main

import (
	"OKProject/api"
	"OKProject/third_party"
	"net/http"
)

func main() {
	// 装载系统配置
	third_party.SetupConfiguration("../")
	mux := http.NewServeMux()
	mux.HandleFunc("/statistics", api.Statistics())
	mux.HandleFunc("/login", api.UserLogin())
	err := http.ListenAndServe(":81", mux)
	if err != nil {
		panic("启动服务失败.")
	}
}
