package config

import (
	"context"
	"fmt"
	log "github.com/sirupsen/logrus"
	"github.com/udayangaac/turbit-nsi/internal/lib/file-manager"
	log_traceable "github.com/udayangaac/turbit-nsi/internal/lib/log-traceable"
	"time"
)

var DatabaseConf DatabaseConfig

type DatabaseConfig struct {
	Host            string        `yaml:"host"`
	Port            int           `yaml:"port"`
	Database        string        `yaml:"database"`
	UserName        string        `yaml:"user_name"`
	Password        string        `yaml:"password"`
	MaxOpenConn     int           `yaml:"max_open_conn"`
	MaxIdleConn     int           `yaml:"max_idle_conn"`
	ConnMaxLifeTime time.Duration `yaml:"conn_max_life_time"`
}

func (sc *DatabaseConfig) Read(fm file_manager.FileManager) {
	path := fmt.Sprintf(`config/database.yaml`)
	err := fm.Read(path, &DatabaseConf)
	if err != nil {
		log.Fatal(log_traceable.GetMessage(context.Background(), "Unable to read the database,yaml file"))
	}
}
