package api

import (
	"OKProject/config"
	"net/http"
)

func Statistics() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("appName:"+config.Config.App.Name))
	}
}
