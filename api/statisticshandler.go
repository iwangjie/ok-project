package api

import (
	"OKProject/config"
	"fmt"
	"net/http"
)

func Statistics() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte("appName:" + config.Config.App.Name))
	}
}

func UserLogin() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		_ = r.ParseForm()
		username := r.Form.Get("username")
		password := r.Form.Get("password")
		fmt.Println("login user,username:"+username, ",password:"+password)
		_, _ = w.Write([]byte("hello. " + username))
	}
}
