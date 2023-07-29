package env

import "github.com/kelseyhightower/envconfig"

// See: https://github.com/l3tnun/EPGStation/blob/master/doc/conf-manual.md

// RecordingCommandEnv コマンドに渡される変数
type RecordingCommandEnv struct {
	RecordedID           string `envconfig:"RECORDEDID" default:"None"`
	ProgramID            string `envconfig:"PROGRAMID" default:"None"`
	ChannelType          string `envconfig:"CHANNELTYPE" default:"None"`
	ChannelID            string `envconfig:"CHANNELID" default:"None"`
	ChannelName          string `envconfig:"CHANNELNAME" default:"None"`
	HalfWidthChannelName string `envconfig:"HALF_WIDTH_CHANNELNAME" default:"None"`
	StartAt              string `envconfig:"STARTAT" default:"None"`
	EndAt                string `envconfig:"ENDAT" default:"None"`
	Duration             string `envconfig:"DURATION" default:"None"`
	Name                 string `envconfig:"NAME" default:"None"`
	HalfWidthName        string `envconfig:"HALF_WIDTH_NAME" default:"None"`
	Description          string `envconfig:"DESCRIPTION" default:"None"`
	HalfWidthDescription string `envconfig:"HALF_WIDTH_DESCRIPTION" default:"None"`
	Extended             string `envconfig:"EXTENDED" default:"None"`
	HalfWidthExtended    string `envconfig:"HALF_WIDTH_EXTENDED" default:"None"`
	RecPath              string `envconfig:"RECPATH" default:"None"`
	LogPath              string `envconfig:"LOGPATH" default:"None"`
}

// EncodingCommandEnv エンコーディング関連のコマンドで渡される環境変数
type EncodingCommandEnv struct {
	RecordedID           string `envconfig:"RECORDEDID" default:"None"`
	VideoFileID          string `envconfig:"VIDEOFILEID" default:"None"`
	OutputPath           string `envconfig:"OUTPUTPATH" default:"None"`
	Mode                 string `envconfig:"MODE" default:"None"`
	ChannelID            string `envconfig:"CHANNELID" default:"None"`
	ChannelName          string `envconfig:"CHANNELNAME" default:"None"`
	HalfWidthChannelName string `envconfig:"HALF_WIDTH_CHANNELNAME" default:"None"`
	Name                 string `envconfig:"NAME" default:"None"`
	HalfWidthName        string `envconfig:"HALF_WIDTH_NAME" default:"None"`
	Description          string `envconfig:"DESCRIPTION" default:"None"`
	HalfWidthDescription string `envconfig:"HALF_WIDTH_DESCRIPTION" default:"None"`
	Extended             string `envconfig:"EXTENDED" default:"None"`
	HalfWidthExtended    string `envconfig:"HALF_WIDTH_EXTENDED" default:"None"`
}

// LoadRecordingCommandEnv 録画関連コマンドで渡される環境変数を読み込む
func LoadRecordingCommandEnv() (*RecordingCommandEnv, error) {
	var env RecordingCommandEnv

	if err := envconfig.Process("", &env); err != nil {
		return nil, err
	}

	return &env, nil
}
