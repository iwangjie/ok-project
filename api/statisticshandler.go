package api

import (
	"OKProject/config"
	"net/http"
)

func Statistics() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte("appName:" + config.Config.App.Name))
	}
}
