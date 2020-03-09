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

func startPreCommandNotification(context *cli.Context, env PreCommandEnv, config Config, commandConfig CommandConfigStruct) error {
	if !commandConfig.Enable {
		displayCommandIsDisableMessage(context)
		return nil
	}

	slackAPIKey := config.Slack.APIKey
	slackChannel := getSlackChannel(config.Slack.Channel, commandConfig.Channel)
	slackClient := createSlackClient(slackAPIKey, context.Bool("debug"))
	message, err := buildPreCommandHeaderText(commandConfig.Message, env)

	if err != nil {
		return err
	}

	blocks := buildPreCommandMessageOptions(message, env)
	_, _, err = slackClient.PostMessage(slackChannel, blocks)

	if err != nil {
		return err
	}

	return nil
}

func startRecCommandNotification(context *cli.Context, env RecCommandEnv, config Config, commandConfig CommandConfigStruct) error {
	if !commandConfig.Enable {
		displayCommandIsDisableMessage(context)
		return nil
	}

	slackAPIKey := config.Slack.APIKey
	slackChannel := getSlackChannel(config.Slack.Channel, commandConfig.Channel)
	slackClient := createSlackClient(slackAPIKey, context.Bool("debug"))
	message, err := buildRecCommandHeaderText(commandConfig.Message, env)

	if err != nil {
		return err
	}

	blocks := buildRecCommandBlocks(message, env)
	_, _, err = slackClient.PostMessage(slackChannel, blocks)

	if err != nil {
		return err
	}

	return nil
}

func displayCommandIsDisableMessage(context *cli.Context) {
	fmt.Printf("%s command is disabled.\n", context.Command.Name)
}
