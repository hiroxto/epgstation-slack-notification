package main

import (
	"github.com/urfave/cli/v2"
	"log"
)

var commandReservationAdded = &cli.Command{
	Name:  "reservation-added",
	Usage: "reservationAddedCommand で使用",
	Description: `
   録画予約の新規追加時に実行されるコマンド.
   reservationAddedCommand で使用する.
`,
	Action: commandReservationAddedAction,
}

func commandReservationAddedAction(context *cli.Context) error {
	config := loadConfigFile()
	env := loadPreCommandEnvs()
	commandConfig := config.Commands.ReservationAdded

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
