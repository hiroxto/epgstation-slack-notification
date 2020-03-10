package main

import (
	"github.com/urfave/cli/v2"
	"log"
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
	config := loadConfigFile()
	env := loadRecCommandEnv()
	err := startRecCommandNotification(context, env, config, config.Commands.RecordedEnd)

	if err != nil {
		log.Fatal(err)
	}

	return nil
}
