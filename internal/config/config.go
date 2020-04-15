package config

import file_manager "github.com/udayangaac/turbit-nsi/internal/lib/file-manager"

func InitConfigurations() {
	fileManager := file_manager.NewYamlManager()
	EnvConf.Read(fileManager)
	DatabaseConf.Read(fileManager)
	ElasticsearchConf.Read(fileManager)
}
