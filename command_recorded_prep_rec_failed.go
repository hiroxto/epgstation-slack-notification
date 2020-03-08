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
	config := loadConfigFile()
	env := loadPreCommandEnvs()
	commandConfig := config.Commands.RecordedPrepRecFailed

	if !commandConfig.Enable {
		displayCommandIsDisableMessage(context)
		return nil
	}

	slackAPIKey := config.Slack.APIKey
	slackChannel := getSlackChannel(config.Slack.Channel, commandConfig.Channel)
	slackClient := createSlackClient(slackAPIKey, context.Bool("debug"))
	blocks := buildPreCommandBlocks(commandConfig.Message, env)
	_, _, err := slackClient.PostMessage(slackChannel, blocks)

	if err != nil {
		log.Fatal(err.Error())
	}

	return nil
}
