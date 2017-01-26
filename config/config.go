package config

import (
	"fmt"
	//"github.com/spf13/viper"
)

type Config struct {
	LogLevel	string
}

func Load() {
	fmt.Println("loading")
}

// Hello returns a dummy text.
func Hello() string {
	return "hello"
}
