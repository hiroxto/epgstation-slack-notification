package main

import (
	"fmt"
	"github.com/urfave/cli/v2"
)

var commands = []*cli.Command{
	commandReservationAdded,
	commandRecordedPreStart,
	commandRecordedPrepRecFailed,

	commandRecordedStart,
	commandRecordedEnd,
	commandRecordedFailed,
}

func startCommandNotification(context *cli.Context, env CommandEnv, config Config, commandConfig CommandConfig) error {
	if !commandConfig.Enable {
		fmt.Printf("%s command is disabled.\n", context.Command.Name)
		return nil
	}

	slackAPIKey := config.Slack.APIKey
	slackChannel := config.Slack.Channel
	if len(commandConfig.Channel) > 0 {
		slackChannel = commandConfig.Channel
	}

	slackClient, err := createSlackClient(slackAPIKey, context.Bool("debug"))
	if err != nil {
		return err
	}

	message, err := formatCommandEnv("", commandConfig.Message, env)

	if err != nil {
		return err
	}

	fields, err := buildCommandFields(commandConfig.FieldsSection, env)

	if err != nil {
		return err
	}

	options := buildMessageOptions(message, fields)
	_, _, err = slackClient.PostMessage(slackChannel, options)

	if err != nil {
		return err
	}

	return nil
}
