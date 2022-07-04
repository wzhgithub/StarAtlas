package db

import (
	"star_atlas_server/config"

	"github.com/golang/glog"
	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Init() error {
	// Setup the mgm default config
	glog.Info("init mongodb start")
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

	_ = mgm.CollectionByName("test_collection").Create(t)

	r := &testDemo{}

	mgm.CollectionByName("test_collection").First(bson.M{"pages": 243242}, r)

	glog.Infof("save and get test %+v", r)

	return nil
}
