package orm

import (
	"context"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	log "github.com/sirupsen/logrus"
	"github.com/udayangaac/turbit-nsi/internal/config"
	log_traceable "github.com/udayangaac/turbit-nsi/internal/lib/log-traceable"
)

var DB *gorm.DB

func InitDatabase(dbConf config.DatabaseConfig) (err error) {
	connectionString := fmt.Sprintf(
		"%v:%v@tcp(%v:%v)/%v?charset=utf8&parseTime=True&loc=Local",
		dbConf.UserName,
		dbConf.Password,
		dbConf.Host,
		dbConf.Port,
		dbConf.Database,
	)
	DB, err = gorm.Open("mysql", connectionString)
	if err != nil {
		return
	}
	DB.SetLogger(&customLogger{})
	DB.LogMode(true)
	return
}

func CloseDatabase() (err error) {
	err = DB.Close()
	return
}

type customLogger struct{}

func (c *customLogger) Print(v ...interface{}) {
	log.Info(log_traceable.GetMessage(context.Background(), v))
}
