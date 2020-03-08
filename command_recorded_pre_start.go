package main

import (
	"github.com/urfave/cli/v2"
	"log"
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
	config := loadConfigFile()
	env := loadPreCommandEnvs()
	commandConfig := config.Commands.RecordedPreStart

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
