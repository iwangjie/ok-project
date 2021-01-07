package config

var Config AppConfig


type AppConfig struct {
	App struct {
		Name string
	}
	Db struct {
		Url string
	}
}
