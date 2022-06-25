package config

import (
	"fmt"

	"github.com/spf13/viper"
)

func init() {
	viper.SetConfigFile("./config.yaml") // 指定配置文件路径
	viper.AddConfigPath(".")             // 设置配置文件和可执行二进制文件在用一个目录

	err := viper.ReadInConfig() // 查找并读取配置文件
	if err != nil {             // 处理读取配置文件的错误
		panic(fmt.Errorf("Fatal error config file: %s \n", err.Error()))
	}
}
