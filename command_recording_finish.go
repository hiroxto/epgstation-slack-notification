package main

import (
	"github.com/urfave/cli/v2"
)

var commandRecordingFinish = &cli.Command{
	Name:    "recording-finish",
	Aliases: []string{"recorded-end"},
	Usage:   "recordingFinishCommand で使用",
	Description: `
   録画終了時に実行するコマンド．
   recordingFinishCommand で使用する．
`,
	Action: commandRecordingFinishAction,
}

func commandRecordingFinishAction(context *cli.Context) error {
	config, err := loadConfigFile()
	if err != nil {
		return err
	}

	var env RecordingCommandEnv
	if err := loadCommandEnv(&env); err != nil {
		return err
	}

	err = startCommandNotification(context, env, config, config.Commands.RecordingFinish)
	if err != nil {
		return err
	}

	err = startRecordedLogNotification(context, env, config, config.Commands.RecordingFinish)
	if err != nil {
		return err
	}

	return nil
}
