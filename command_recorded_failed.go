package main

import (
	"github.com/urfave/cli/v2"
	"log"
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
	config := loadConfigFile()
	env := loadRecCommandEnv()
	err := startRecCommandNotification(context, env, config, config.Commands.RecordedFailed)

	if err != nil {
		log.Fatal(err.Error())
	}

	return nil
}
