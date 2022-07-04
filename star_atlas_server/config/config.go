package config

import (
	"github.com/golang/glog"
	"github.com/spf13/viper"
)

type Config struct {
	DB                 string `mapstructure:"db_name"`
	Uri                string `mapstructure:"db_uri"`
	DBVMCDataTableName string `mapstructure:"db_vmcdata_table_name"`
	UDPPort            int    `mapstructure:"udp_port"`
}

var CommonConfig Config

func Init(path string) error {
	glog.Infof("path:%s\n", path)
	viper.SetConfigFile(path)
	// 查找并读取配置文件
	if err := viper.ReadInConfig(); err != nil {
		return err
	}

	if err := viper.Unmarshal(&CommonConfig); err != nil {
		return err
	}
	return nil
}
