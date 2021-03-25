package main

import (
	"github.com/urfave/cli/v2"
)

var commandRecordingPrepRecFailed = &cli.Command{
	Name:    "recording-prep-rec-failed",
	Aliases: []string{"recorded-prep-rec-failed"},
	Usage:   "recordingPrepRecFailedCommand で使用",
	Description: `
   録画準備の失敗時に実行されるコマンド．
   recordingPrepRecFailedCommand で使用する．`,
	Action: commandRecordingPrepRecFailedAction,
}

func commandRecordingPrepRecFailedAction(context *cli.Context) error {
	config, err := loadConfigFile()
	if err != nil {
		return err
	}

	env, err := loadCommandEnv()
	if err != nil {
		return err
	}

	err = startCommandNotification(context, env, config, config.Commands.RecordedPrepRecFailed)
	if err != nil {
		return err
	}

	return nil
}
