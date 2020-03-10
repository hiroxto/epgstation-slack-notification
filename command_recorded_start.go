package main

import (
	"github.com/urfave/cli/v2"
	"log"
)

var commandRecordedStart = &cli.Command{
	Name:  "recorded-start",
	Usage: "recordedStartCommand で使用",
	Description: `
   録画予約の新規追加時に実行されるコマンド.
   recordedStartCommand で使用する.
`,
	Action: commandRecordedStartAction,
}

func commandRecordedStartAction(context *cli.Context) error {
	config, err := loadConfigFile()

	if err != nil {
		log.Fatal(err)
	}

	env, err := loadRecCommandEnv()
	if err != nil {
		log.Fatal(err)
	}

	err = startRecCommandNotification(context, env, config, config.Commands.RecordedStart)

	if err != nil {
		log.Fatal(err)
	}

	return nil
}
