package main

import (
	"github.com/urfave/cli/v2"
)

var commandRecordedFailed = &cli.Command{
	Name:  "recorded-failed",
	Usage: "recordedFailedCommand で使用",
	Description: `
   録画中のエラー発生時に実行するコマンド.
   recordedFailedCommand で使用する.
`,
	Action: commandRecordedFailedAction,
}

func commandRecordedFailedAction(context *cli.Context) error {
	config, err := loadConfigFile()
	if err != nil {
		return err
	}

	env, err := loadCommandEnv()
	if err != nil {
		return err
	}

	err = startCommandNotification(context, env, config, config.Commands.RecordedFailed)
	if err != nil {
		return err
	}

	return nil
}
