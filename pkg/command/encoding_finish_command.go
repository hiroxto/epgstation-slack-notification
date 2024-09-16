package command

import (
	"fmt"

	"github.com/hiroxto/epgstation-slack-notification/pkg/app"
	"github.com/hiroxto/epgstation-slack-notification/pkg/config"
	"github.com/hiroxto/epgstation-slack-notification/pkg/env"
	"github.com/urfave/cli/v2"
)

var EncodingFinishCommand = &cli.Command{
	Name:  "encoding-finish",
	Usage: "encodingFinishCommand で使用",
	Description: `
   エンコード終了時に実行するコマンド。
   encodingFinishCommand で使用する。
`,
	Action: encodingFinishCommandAction,
}

func encodingFinishCommandAction(context *cli.Context) error {
	configFilePath := getConfigFilePath(context)
	configData, err := readConfigFile(configFilePath)
	if err != nil {
		return err
	}

	conf, err := config.LoadConfigFromYaml([]byte(configData))
	if err != nil {
		return err
	}

	commandConfig := conf.Commands.EncodingFinish
	if !commandConfig.Enable {
		fmt.Printf("%s command is disabled.\n", context.Command.Name)
		return nil
	}
	slackAPIKey := conf.Slack.APIKey
	slackChannel := conf.Slack.Channel
	if len(commandConfig.Channel) > 0 {
		slackChannel = commandConfig.Channel
	}

	encodingCommandEnv, err := env.LoadEncodingCommandEnv()
	if err != nil {
		return err
	}

	param := app.EncodingUseCaseParam{
		EnableDebug:    context.Bool("debug"),
		SlackAPIKey:    slackAPIKey,
		SlackChannel:   slackChannel,
		UserName:       commandConfig.UserName,
		Message:        commandConfig.Message,
		Fields:         app.FieldsFromConfig(commandConfig.FieldsSection),
		EncodingDetail: app.EncodingDetailFromEnv(encodingCommandEnv),
	}
	err = app.EncodingNotificationUseCase(param)
	if err != nil {
		return err
	}

	return nil
}
