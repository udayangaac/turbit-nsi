package config

import (
	"context"
	log "github.com/sirupsen/logrus"
	file_manager "github.com/udayangaac/turbit-nsi/internal/lib/file-manager"
	log_traceable "github.com/udayangaac/turbit-nsi/internal/lib/log-traceable"
)

var EnvConf EnvConfig

type EnvConfig struct {
	Port     int `yaml:"port"`
	H3Config struct {
		Resolution int `yaml:"resolution"`
	} `yaml:"h3_config"`
	Badger struct {
		DbPath string `yaml:"db_path"`
	} `yaml:"badger"`
}

func (sc *EnvConfig) Read(fm file_manager.FileManager) {
	err := fm.Read(`config/env.yaml`, &EnvConf)
	if err != nil {
		log.Fatal(log_traceable.GetMessage(context.Background(), "Unable to read the database,yaml file"))
	}
}
