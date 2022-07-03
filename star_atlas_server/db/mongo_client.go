package db

import (
	"star_atlas_server/config"

	"github.com/golang/glog"
	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Init() error {
	// Setup the mgm default config
	glog.Info("init mongodb init")
	err := mgm.SetDefaultConfig(nil, config.CommonConfig.DB, options.Client().ApplyURI(config.CommonConfig.Uri))
	if err != nil {
		return err
	}

	return nil
}

// see https://github.com/Kamva/mgm
type testDemo struct {
	mgm.DefaultModel `bson:",inline"`
	Name             string `json:"name" bson:"name"`
	Pages            int    `json:"pages" bson:"pages"`
}

func Test() error {
	t := &testDemo{
		Name:  "xxxx",
		Pages: 243242,
	}

	return mgm.CollectionByName("test_collection").Create(t)
}
