package main

import (
	"fmt"
	"time"

	"github.com/spf13/viper"
)

func main() {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	viper.SetConfigType("yaml")

	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("Error reading config file, %s", err)
		return
	}

	// Theo dõi và nạp lại cấu hình khi có thay đổi
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("Config file changed:", e.Name)
	})

	for {
		appName := viper.GetString("app_name")
		fmt.Println("App Name:", appName)
		time.Sleep(5 * time.Second)
	}
}
