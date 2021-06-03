package main

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type config struct {
	Port string       `yaml:"Port"`
	Folders []string  `yaml:"Folders"`
}

func (config *config) fillFromYML(ymlFileName string) error {
	content, err := ioutil.ReadFile(ymlFileName)
	if err != nil {
		return err
	}

	err = yaml.Unmarshal(content, &config)
	if err != nil {
		return err
	}

	return nil
}