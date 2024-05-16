package command

import (
	"fmt"

	"github.com/hiroxto/epgstation-slack-notification/pkg/app"
	"github.com/hiroxto/epgstation-slack-notification/pkg/config"
	"github.com/hiroxto/epgstation-slack-notification/pkg/env"
	"github.com/urfave/cli/v2"
)

var RecordingStartCommand = &cli.Command{
	Name:    "recording-start",
	Aliases: []string{"recorded-start"},
	Usage:   "recordingStartCommand で使用",
	Description: `
   録画予約の新規追加時に実行されるコマンド．
   recordingStartCommand で使用する．
`,
	Action: recordingStartCommandAction,
}

func recordingStartCommandAction(context *cli.Context) error {
	configFilePath := getConfigFilePath(context)
	configData, err := readConfigFile(configFilePath)
	if err != nil {
		return err
	}

	config, err := config.LoadConfigFromYaml([]byte(configData))
	if err != nil {
		return err
	}

	commandConfig := config.Commands.RecordingStart
	if !commandConfig.Enable {
		fmt.Printf("%s command is disabled.\n", context.Command.Name)
		return nil
	}
	slackAPIKey := config.Slack.APIKey
	slackChannel := config.Slack.Channel
	if len(commandConfig.Channel) > 0 {
		slackChannel = commandConfig.Channel
	}

	recordingCommandEnv, err := env.LoadRecordingCommandEnv()
	if err != nil {
		return err
	}

	param := app.RecordingUseCaseParam{
		EnableDebug:     context.Bool("debug"),
		SlackAPIKey:     slackAPIKey,
		SlackChannel:    slackChannel,
		Message:         commandConfig.Message,
		Fields:          app.FieldsFromConfig(commandConfig.FieldsSection),
		RecordingDetail: app.RecordingDetailFromEnv(recordingCommandEnv),
	}
	err = app.RecordingNotificationUseCase(param)
	if err != nil {
		return err
	}

	return nil
}
