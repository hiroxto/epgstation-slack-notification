package main

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
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
		ReservationAdded      CommandConfig `yaml:"reservation-added"`
		RecordedPreStart      CommandConfig `yaml:"recorded-pre-start"`
		RecordedPrepRecFailed CommandConfig `yaml:"recorded-prep-rec-failed"`
		RecordedStart         CommandConfig `yaml:"recorded-start"`
		RecordedEnd           CommandConfig `yaml:"recorded-end"`
		RecordedFailed        CommandConfig `yaml:"recorded-failed"`
	} `yaml:"commands"`
}

// CommandConfig 各コマンドの設定
type CommandConfig struct {
	Enable        bool            `yaml:"enable"`
	Channel       string          `yaml:"channel"`
	Message       string          `yaml:"message"`
	FieldsSection []FieldsSection `yaml:"fields-section"`
}

// FieldsSection Slack の fields の設定
type FieldsSection struct {
	Title    string `yaml:"title"`
	Template string `yaml:"template"`
}

func loadConfigFile() (Config, error) {
	var config Config

	configFilePath, err := getConfigFilePath()
	if err != nil {
		return config, err
	}

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

func getConfigFilePath() (string, error) {
	exeFilePath, err := os.Executable()
	if err != nil {
		return exeFilePath, err
	}

	return filepath.Join(filepath.Dir(exeFilePath), "epgstation-slack-config.yml"), nil
}
