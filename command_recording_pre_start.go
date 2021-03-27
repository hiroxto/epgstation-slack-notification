package main

import (
	"github.com/urfave/cli/v2"
)

var commandRecordingPreStart = &cli.Command{
	Name:    "recording-pre-start",
	Aliases: []string{"recorded-pre-start"},
	Usage:   "recordingPreStartCommand で使用",
	Description: `
   録画準備の開始時に実行されるコマンド．
   recordingPreStartCommand で使用する．
`,
	Action: commandRecordingPreStartAction,
}

func commandRecordingPreStartAction(context *cli.Context) error {
	config, err := loadConfigFile()
	if err != nil {
		return err
	}

	env, err := loadCommandEnv()
	if err != nil {
		return err
	}

	err = startCommandNotification(context, env, config, config.Commands.RecordingPreStart)
	if err != nil {
		return err
	}

	return nil
}
