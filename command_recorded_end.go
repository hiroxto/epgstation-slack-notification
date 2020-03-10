package main

import (
	"github.com/urfave/cli/v2"
)

var commandRecordedEnd = &cli.Command{
	Name:  "recorded-end",
	Usage: "recordedEndCommand で使用",
	Description: `
   録画終了時に実行するコマンド.
   recordedEndCommand で使用する.
`,
	Action: commandRecordedEndAction,
}

func commandRecordedEndAction(context *cli.Context) error {
	config, err := loadConfigFile()
	if err != nil {
		return err
	}

	env, err := loadRecCommandEnv()
	if err != nil {
		return err
	}

	err = startRecCommandNotification(context, env, config, config.Commands.RecordedEnd)
	if err != nil {
		return err
	}

	return nil
}
