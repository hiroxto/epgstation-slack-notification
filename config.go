package main

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

// Config epg-slack-config.yml で渡される設定
type Config struct {
	Slack struct {
		APIKey  string `yaml:"api-key"`
		Channel string `yaml:"channel"`
	} `yaml:"slack"`
	Commands struct {
		ReservationAdded      CommandConfigStruct `yaml:"reservation-added"`
		RecordedPreStart      CommandConfigStruct `yaml:"recorded-pre-start"`
		RecordedPrepRecFailed CommandConfigStruct `yaml:"recorded-prep-rec-failed"`
		RecordedStart         CommandConfigStruct `yaml:"recorded-start"`
		RecordedEnd           CommandConfigStruct `yaml:"recorded-end"`
		RecordedFailed        CommandConfigStruct `yaml:"recorded-failed"`
	} `yaml:"commands"`
}

// CommandConfigStruct 各コマンドの設定
type CommandConfigStruct struct {
	Enable        bool                  `yaml:"enable"`
	Channel       string                `yaml:"channel"`
	Message       string                `yaml:"message"`
	FieldsSection []FieldsSectionStruct `yaml:"fields-section"`
}

// FieldsSectionStruct Slack の fields の設定
type FieldsSectionStruct struct {
	Title    string `yaml:"title"`
	Template string `yaml:"template"`
}

func loadConfigFile() (Config, error) {
	var config Config
	configFilePath := getConfigFilePath()

	data, err := ioutil.ReadFile(configFilePath)
	if err != nil {
		return config, err
	}

	err = yaml.UnmarshalStrict([]byte(data), &config)
	if err != nil {
		return config, err
	}

	return config, nil
}

func getConfigFilePath() string {
	exeFilePath, err := os.Executable()
	if err != nil {
		log.Fatal(err.Error())
	}

	return filepath.Join(filepath.Dir(exeFilePath), "epgstation-slack-config.yml")
}
