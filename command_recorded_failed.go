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
	commandConfig := config.Commands.RecordedFailed

	if !commandConfig.Enable {
		displayCommandIsDisableMessage(context)
		return nil
	}

	slackAPIKey := config.Slack.APIKey
	slackChannel := getSlackChannel(config.Slack.Channel, commandConfig.Channel)
	slackClient := createSlackClient(slackAPIKey, context.Bool("debug"))
	blocks := buildRecCommandBlocks(commandConfig.Message, env)
	_, _, err := slackClient.PostMessage(slackChannel, blocks)

	if err != nil {
		log.Fatal(err.Error())
	}

	return nil
}
