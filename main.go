package main

import (
	"OKProject/api"
	"OKProject/third_party"
	"net/http"
)

type MyHandler struct {
}

func (m MyHandler) ServeHTTP(writer http.ResponseWriter, r *http.Request) {
	_, _ = writer.Write([]byte("测试一下~~~......"))
}

func main() {
	// 装载系统配置
	third_party.SetupConfiguration("../")
	mux := http.NewServeMux()
	mux.HandleFunc("/statistics", api.Statistics())
	err := http.ListenAndServe(":81", mux)
	if err != nil {
		panic("启动服务失败.")
	}
}
