package main

import (
	"github.com/urfave/cli/v2"
)

var commandRecordingFailed = &cli.Command{
	Name:    "recording-failed",
	Aliases: []string{"recorded-failed"},
	Usage:   "recordingFailedCommand で使用",
	Description: `
   録画中のエラー発生時に実行するコマンド．
   recordingFailedCommand で使用する．
`,
	Action: commandRecordingFailedAction,
}

func commandRecordingFailedAction(context *cli.Context) error {
	config, err := loadConfigFile()
	if err != nil {
		return err
	}

	var env CommandEnv
	if err := loadCommandEnv(&env); err != nil {
		return err
	}

	err = startCommandNotification(context, env, config, config.Commands.RecordingFailed)
	if err != nil {
		return err
	}

	return nil
}
