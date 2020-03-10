package main

import (
	"github.com/urfave/cli/v2"
)

var commandRecordedPreStart = &cli.Command{
	Name:  "recorded-pre-start",
	Usage: "recordedPreStartCommand で使用",
	Description: `
   録画準備の開始時に実行されるコマンド.
   recordedPreStartCommand で使用する.
`,
	Action: commandRecordedPreStartAction,
}

func commandRecordedPreStartAction(context *cli.Context) error {
	config, err := loadConfigFile()
	if err != nil {
		return err
	}

	env, err := loadPreCommandEnv()
	if err != nil {
		return err
	}

	err = startPreCommandNotification(context, env, config, config.Commands.RecordedPreStart)
	if err != nil {
		return err
	}

	return nil
}
