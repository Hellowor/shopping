package conf

import (
	"github.com/jinzhu/configor"
	"github.com/rex-ss/library/database/mysql"
	log "github.com/sirupsen/logrus"
	"os"
)

var Config = struct {
	APPName    string `default:"mtActivity"`
	APPPort    string
	IotUrl     string
	SchemeHost string
	TmpFolder  string
	Pwd        string
	DB         *mysql.Config
	DBDAO      *mysql.Dao
}{}

func init() {
	path, err := os.Getwd()
	log.Info("conf.init(), pwd:", path)
	//err = configor.Load(&Config, path+"/src/mtactivity/conf/conf.yaml")
	err = configor.Load(&Config, path+"/conf/conf.yaml")
	if err != nil {
		panic(err)
	}
	Config.Pwd = path
	log.Info("config: %#v", Config)
}

func init() {
	Config.DBDAO = mysql.New(Config.DB)
}
