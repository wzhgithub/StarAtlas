package db

import (
	"fmt"
	"log"
	_ "start_atlas_server/config"

	"github.com/kamva/mgm/v3"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func init() {
	// Setup the mgm default config
	log.Println("init mongodb init")
	err := mgm.SetDefaultConfig(nil, viper.GetString("db"), options.Client().ApplyURI(viper.GetString("uri")))
	if err != nil {
		fmt.Println("mgn init failed", err.Error())
	}
}
