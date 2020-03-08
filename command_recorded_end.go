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
	commandConfig := config.Commands.RecordedEnd

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
