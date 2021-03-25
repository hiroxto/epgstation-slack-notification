package main

import (
	"github.com/kelseyhightower/envconfig"
)

// See: https://github.com/l3tnun/EPGStation/blob/master/doc/conf-manual.md

// CommandEnv コマンドに渡される変数
type CommandEnv struct {
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
	Extended             string `envconfig:"EXTENDED" default:"None"`
	RecPath              string `envconfig:"RECPATH" default:"None"`
	LogPath              string `envconfig:"LOGPATH" default:"None"`
}

func loadCommandEnv() (CommandEnv, error) {
	var env CommandEnv

	if err := envconfig.Process("", &env); err != nil {
		return env, err
	}

	return env, nil
}
