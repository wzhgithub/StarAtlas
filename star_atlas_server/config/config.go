package config

import (
	"github.com/golang/glog"
	"github.com/spf13/viper"
)

type Config struct {
	DB                 string `mapstructure:"db_name"`
	Uri                string `mapstructure:"db_uri"`
	UDPPort            int    `mapstructure:"udp_port"`
	SatelliteTCPPort   int    `mapstructure:"satellite_tcp_port"`
	SpeechURL          string `mapstructure:"speech_url"`
	DBVMCDataTableName string `mapstructure:"db_vmcdata_table_name"`
	DBTopoTableName    string `mapstructure:"db_topo_table_name"`
	DBSenderTableName  string `mapstructure:"db_sender_table_name"`
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
