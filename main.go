package main

import (
	"OKProject/api"
	"OKProject/third_party"
	"net/http"
)

type MyHandler struct {
}

func (m MyHandler) ServeHTTP(writer http.ResponseWriter, r *http.Request) {
	writer.Write([]byte("完美~~~......"))
}

func main() {
	// 装载系统配置
	third_party.SetupConfiguration("../")
	mux := http.NewServeMux()
	mux.HandleFunc("/statistics", api.Statistics())
	http.ListenAndServe(":81", mux)
}
