package main

import (
	"github.com/urfave/cli/v2"
)

var commandRecordingStart = &cli.Command{
	Name:    "recording-start",
	Aliases: []string{"recorded-start"},
	Usage:   "recordingStartCommand で使用",
	Description: `
   録画予約の新規追加時に実行されるコマンド．
   recordingStartCommand で使用する．
`,
	Action: commandRecordingStartAction,
}

func commandRecordingStartAction(context *cli.Context) error {
	config, err := loadConfigFile()
	if err != nil {
		return err
	}

	var env CommandEnv
	if err := loadCommandEnv(&env); err != nil {
		return err
	}

	err = startCommandNotification(context, env, config, config.Commands.RecordingStart)
	if err != nil {
		return err
	}

	return nil
}
