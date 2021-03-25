package main

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"path/filepath"
)

// Config epg-slack-config.yml で渡される設定
type Config struct {
	EPGStation struct {
		HostName string `yaml:"host-name"`
	} `yaml:"epg-station"`
	Slack struct {
		APIKey  string `yaml:"api-key"`
		Channel string `yaml:"channel"`
	} `yaml:"slack"`
	Commands struct {
		ReserveNewAddition     CommandConfig `yaml:"reserve-new-addition"`
		ReserveUpdateCommand   CommandConfig `yaml:"reserve-update"`
		ReserveDeletedCommand  CommandEnv    `yaml:"reserve-deleted"`
		RecordingPreStart      CommandConfig `yaml:"recording-pre-start"`
		RecordingPrepRecFailed CommandConfig `yaml:"recording-prep-rec-failed"`
		RecordingStart         CommandConfig `yaml:"recording-start"`
		RecordingFinish        CommandConfig `yaml:"recording-finish"`
		RecordingFailed        CommandConfig `yaml:"recording-failed"`
	} `yaml:"commands"`
}

// CommandConfig 各コマンドの設定
type CommandConfig struct {
	Enable        bool          `yaml:"enable"`
	Channel       string        `yaml:"channel"`
	Message       string        `yaml:"message"`
	FieldsSection []FieldConfig `yaml:"fields-section"`
}

// FieldConfig Slack の fields の設定
type FieldConfig struct {
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
