package main

import (
	"github.com/urfave/cli/v2"
	"log"
)

var commandRecordedPrepRecFailed = &cli.Command{
	Name:  "recorded-prep-rec-failed",
	Usage: "recordedPrepRecFailedCommand で使用",
	Description: `
   録画準備の失敗時に実行されるコマンド.
   recordedPrepRecFailedCommand で使用する.`,
	Action: commandRecordedPrepRecFailedAction,
}

func commandRecordedPrepRecFailedAction(context *cli.Context) error {
	config, err := loadConfigFile()

	if err != nil {
		log.Fatal(err)
	}

	env := loadPreCommandEnvs()
	err = startPreCommandNotification(context, env, config, config.Commands.RecordedPrepRecFailed)

	if err != nil {
		log.Fatal(err)
	}

	return nil
}
