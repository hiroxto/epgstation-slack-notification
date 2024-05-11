package command

import (
	"fmt"

	"github.com/hiroxto/epgstation-slack-notification/pkg/app"
	"github.com/hiroxto/epgstation-slack-notification/pkg/config"
	"github.com/hiroxto/epgstation-slack-notification/pkg/env"
	"github.com/urfave/cli/v2"
)

var CommandEncodingFinish = &cli.Command{
	Name:  "encoding-finish",
	Usage: "encodingFinishCommand で使用",
	Description: `
   エンコード終了時に実行するコマンド。
   encodingFinishCommand で使用する。
`,
	Action: commandEncodingFinishAction,
}

func commandEncodingFinishAction(context *cli.Context) error {
	configFilePath, err := getConfigFilePath()
	if err != nil {
		return err
	}

	configData, err := readConfigFile(configFilePath)
	if err != nil {
		return err
	}

	config, err := config.LoadConfigFromYaml([]byte(configData))
	if err != nil {
		return err
	}

	commandConfig := config.Commands.EncodingFinish
	if !commandConfig.Enable {
		fmt.Printf("%s command is disabled.\n", context.Command.Name)
		return nil
	}
	slackAPIKey := config.Slack.APIKey
	slackChannel := config.Slack.Channel
	if len(commandConfig.Channel) > 0 {
		slackChannel = commandConfig.Channel
	}

	encodingCommandEnv, err := env.LoadEncodingCommandEnv()
	if err != nil {
		return err
	}

	encodingDetail := app.EncodingDetailFromEnv(encodingCommandEnv)
	param := app.EncodingUseCaseParam{
		EnableDebug:    context.Bool("debug"),
		SlackAPIKey:    slackAPIKey,
		SlackChannel:   slackChannel,
		Message:        commandConfig.Message,
		Fields:         []app.Field{},
		EncodingDetail: encodingDetail,
	}
	err = app.EncodingNotificationUseCase(param)
	if err != nil {
		return err
	}

	return nil
}
