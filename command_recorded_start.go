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
	config := loadConfigFile()
	env := loadRecCommandEnv()
	commandConfig := config.Commands.RecordedStart

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
