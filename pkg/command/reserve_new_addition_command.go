package command

import (
	"fmt"

	"github.com/hiroxto/epgstation-slack-notification/pkg/app"
	"github.com/hiroxto/epgstation-slack-notification/pkg/config"
	"github.com/hiroxto/epgstation-slack-notification/pkg/env"
	"github.com/urfave/cli/v2"
)

var ReserveNewAdditionCommand = &cli.Command{
	Name:    "reserve-new-addition",
	Aliases: []string{"reservation-added"},
	Usage:   "reserveNewAddtionCommand で使用",
	Description: `
   録画予約の新規追加時に実行されるコマンド．
   reserveNewAddtionCommand で使用する．
`,
	Action: reserveNewAdditionCommandAction,
}

func reserveNewAdditionCommandAction(context *cli.Context) error {
	configFilePath := getConfigFilePath(context)
	configData, err := readConfigFile(configFilePath)
	if err != nil {
		return err
	}

	conf, err := config.LoadConfigFromYaml([]byte(configData))
	if err != nil {
		return err
	}

	commandConfig := conf.Commands.ReserveNewAddition
	if !commandConfig.Enable {
		fmt.Printf("%s command is disabled.\n", context.Command.Name)
		return nil
	}
	slackAPIKey := conf.Slack.APIKey
	slackChannel := conf.Slack.Channel
	if len(commandConfig.Channel) > 0 {
		slackChannel = commandConfig.Channel
	}

	reserveCommandEnv, err := env.LoadReserveCommandEnv()
	if err != nil {
		return err
	}

	param := app.ReserveUseCaseParam{
		EnableDebug:   context.Bool("debug"),
		SlackAPIKey:   slackAPIKey,
		SlackChannel:  slackChannel,
		UserName:      commandConfig.UserName,
		Message:       commandConfig.Message,
		Fields:        app.FieldsFromConfig(commandConfig.FieldsSection),
		ReserveDetail: app.ReserveDetailFromEnv(reserveCommandEnv),
	}
	err = app.ReserveNotificationUseCase(param)
	if err != nil {
		return err
	}

	return nil
}
