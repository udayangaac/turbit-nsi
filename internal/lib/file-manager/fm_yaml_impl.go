package file_manager

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
)

type yamlManager struct{}

func NewYamlManager() FileManager {
	return &yamlManager{}
}

func (ym *yamlManager) Read(path string, i interface{}) (readErr error) {
	yamlFile, err := ioutil.ReadFile(fmt.Sprintf("./%v", path))
	if err != nil {
		readErr = err
		return
	}
	err = yaml.Unmarshal(yamlFile, i)
	if err != nil {
		readErr = err
		return
	}
	return
}

func (ym *yamlManager) Write(path string, i interface{}) (writeErr error) {
	yamlData, err := yaml.Marshal(i)
	if err != nil {
		writeErr = err
		return
	}
	err = ioutil.WriteFile(fmt.Sprintf("./%v", path), yamlData, os.ModeAppend)
	if err != nil {
		writeErr = err
		return
	}
	return
}
