package config

import "gopkg.in/yaml.v2"

// Config アプリの設定
type Config struct {
	EPGStation EPGStationConfig `yaml:"epg-station"`
	Slack      SlackConfig      `yaml:"slack"`
	Commands   CommandsConfig   `yaml:"commands"`
}

// EPGStationConfig EPGStation環境の設定
type EPGStationConfig struct {
	HostName string `yaml:"host-name"`
}

// SlackConfig Slackの設定
type SlackConfig struct {
	APIKey  string `yaml:"api-key"`
	Channel string `yaml:"channel"`
}

// CommandsConfig コマンドの設定
type CommandsConfig struct {
	ReserveNewAddition     CommandConfig `yaml:"reserve-new-addition"`
	ReserveUpdateCommand   CommandConfig `yaml:"reserve-update"`
	ReserveDeletedCommand  CommandConfig `yaml:"reserve-deleted"`
	RecordingPreStart      CommandConfig `yaml:"recording-pre-start"`
	RecordingPrepRecFailed CommandConfig `yaml:"recording-prep-rec-failed"`
	RecordingStart         CommandConfig `yaml:"recording-start"`
	RecordingFinish        CommandConfig `yaml:"recording-finish"`
	RecordingFailed        CommandConfig `yaml:"recording-failed"`
	EncodingFinish         CommandConfig `yaml:"encoding-finish"`
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

// LoadConfigFromYaml YAMLから設定を読み込む
func LoadConfigFromYaml(configYaml []byte) (Config, error) {
	var config Config

	err := yaml.UnmarshalStrict(configYaml, &config)
	if err != nil {
		return config, err
	}

	return config, nil
}
