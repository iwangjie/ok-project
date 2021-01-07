package third_party

import (
	"OKProject/config"
	"fmt"
	"github.com/spf13/viper"
)



func SetupConfiguration(configPath string) {
	viper.SetConfigName("config") // name of config file (without extension)
	viper.SetConfigType("yaml") // REQUIRED if the config file does not have the extension in the name
	viper.AddConfigPath("/etc/ok_project/")   // path to look for the config file in
	viper.AddConfigPath("$HOME/.ok_project")  // call multiple times to add many search paths
	viper.AddConfigPath(".")               // optionally look for config in the working directory
	viper.AddConfigPath(configPath)
	err := viper.ReadInConfig() // Find and read the config file
	if err != nil { // Handle errors reading the config file
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	viper.Unmarshal(&config.Config)
}
