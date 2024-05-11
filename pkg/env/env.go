package env

import "github.com/kelseyhightower/envconfig"

// See: https://github.com/l3tnun/EPGStation/blob/master/doc/conf-manual.md

// ReserveCommandEnv 予約コマンドと録画準備コマンドに渡される環境変数
type ReserveCommandEnv struct {
	ProgramID            int    `envconfig:"PROGRAMID"`
	ChannelType          string `envconfig:"CHANNELTYPE" default:"None"`
	ChannelID            int    `envconfig:"CHANNELID"`
	ChannelName          string `envconfig:"CHANNELNAME" default:"None"`
	HalfWidthChannelName string `envconfig:"HALF_WIDTH_CHANNELNAME" default:"None"`
	StartAt              int    `envconfig:"STARTAT"`
	EndAt                int    `envconfig:"ENDAT"`
	Duration             int    `envconfig:"DURATION"`
	Name                 string `envconfig:"NAME" default:"None"`
	HalfWidthName        string `envconfig:"HALF_WIDTH_NAME" default:"None"`
	Description          string `envconfig:"DESCRIPTION" default:"None"`
	HalfWidthDescription string `envconfig:"HALF_WIDTH_DESCRIPTION" default:"None"`
	Extended             string `envconfig:"EXTENDED" default:"None"`
	HalfWidthExtended    string `envconfig:"HALF_WIDTH_EXTENDED" default:"None"`
}

// RecordingCommandEnv 録画コマンドに渡される変数
type RecordingCommandEnv struct {
	RecordedID           int    `envconfig:"RECORDEDID"`
	ProgramID            int    `envconfig:"PROGRAMID"`
	ChannelType          string `envconfig:"CHANNELTYPE" default:"None"`
	ChannelID            int    `envconfig:"CHANNELID"`
	ChannelName          string `envconfig:"CHANNELNAME" default:"None"`
	HalfWidthChannelName string `envconfig:"HALF_WIDTH_CHANNELNAME" default:"None"`
	StartAt              int    `envconfig:"STARTAT"`
	EndAt                int    `envconfig:"ENDAT"`
	Duration             int    `envconfig:"DURATION"`
	Name                 string `envconfig:"NAME" default:"None"`
	HalfWidthName        string `envconfig:"HALF_WIDTH_NAME" default:"None"`
	Description          string `envconfig:"DESCRIPTION" default:"None"`
	HalfWidthDescription string `envconfig:"HALF_WIDTH_DESCRIPTION" default:"None"`
	Extended             string `envconfig:"EXTENDED" default:"None"`
	HalfWidthExtended    string `envconfig:"HALF_WIDTH_EXTENDED" default:"None"`
	RecPath              string `envconfig:"RECPATH" default:"None"`
	LogPath              string `envconfig:"LOGPATH" default:"None"`
	ErrorCnt             int    `envconfig:"ERROR_CNT"`
	DropCnt              int    `envconfig:"DROP_CNT"`
	ScramblingCount      int    `envconfig:"SCRAMBLING_CNT"`
}

// EncodingCommandEnv エンコードコマンドに渡される環境変数
type EncodingCommandEnv struct {
	RecordedID           int    `envconfig:"RECORDEDID"`
	VideoFileID          int    `envconfig:"VIDEOFILEID"`
	OutputPath           string `envconfig:"OUTPUTPATH" default:"None"`
	Mode                 string `envconfig:"MODE" default:"None"`
	ChannelID            int    `envconfig:"CHANNELID"`
	ChannelName          string `envconfig:"CHANNELNAME" default:"None"`
	HalfWidthChannelName string `envconfig:"HALF_WIDTH_CHANNELNAME" default:"None"`
	Name                 string `envconfig:"NAME" default:"None"`
	HalfWidthName        string `envconfig:"HALF_WIDTH_NAME" default:"None"`
	Description          string `envconfig:"DESCRIPTION" default:"None"`
	HalfWidthDescription string `envconfig:"HALF_WIDTH_DESCRIPTION" default:"None"`
	Extended             string `envconfig:"EXTENDED" default:"None"`
	HalfWidthExtended    string `envconfig:"HALF_WIDTH_EXTENDED" default:"None"`
}

// LoadReserveCommandEnv 予約コマンドと録画準備コマンドで渡される環境変数を読み込む
func LoadReserveCommandEnv() (ReserveCommandEnv, error) {
	var env ReserveCommandEnv

	if err := envconfig.Process("", &env); err != nil {
		return env, err
	}

	return env, nil
}

// LoadRecordingCommandEnv 録画コマンドで渡される環境変数を読み込む
func LoadRecordingCommandEnv() (RecordingCommandEnv, error) {
	var env RecordingCommandEnv

	if err := envconfig.Process("", &env); err != nil {
		return env, err
	}

	return env, nil
}

// LoadEncodingCommandEnv エンコードコマンドで渡される環境変数を読み込む
func LoadEncodingCommandEnv() (EncodingCommandEnv, error) {
	var env EncodingCommandEnv

	if err := envconfig.Process("", &env); err != nil {
		return env, err
	}

	return env, nil
}
