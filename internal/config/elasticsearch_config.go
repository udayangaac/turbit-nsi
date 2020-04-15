// Copyright 2020. All rights reserved.
// Author : Chamith Udayanga.
// Use of this source code is governed by a
// license that can be found in the LICENSE file.

package config

import (
	"context"
	"fmt"
	log "github.com/sirupsen/logrus"
	"github.com/udayangaac/turbit-nsi/internal/lib/file-manager"
	log_traceable "github.com/udayangaac/turbit-nsi/internal/lib/log-traceable"
)

var ElasticsearchConf ElasticsearchConfig

type ElasticsearchConfig struct {
	Addresses []string `yaml:"addresses"`
	Username  string   `yaml:"username"`
	Password  string   `yaml:"password"`
}

func (sc *ElasticsearchConfig) Read(fm file_manager.FileManager) {
	path := fmt.Sprintf(`config/elasticsearch.yaml`)
	err := fm.Read(path, &ElasticsearchConf)
	if err != nil {
		log.Fatal(log_traceable.GetMessage(context.Background(), "Unable to read the elasticsearch.yaml file."))
	}
}
