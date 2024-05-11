package command

import (
	"fmt"

	"github.com/hiroxto/epgstation-slack-notification/pkg/app"
	"github.com/hiroxto/epgstation-slack-notification/pkg/config"
	"github.com/hiroxto/epgstation-slack-notification/pkg/env"
	"github.com/urfave/cli/v2"
)

var RecordingFailedCommand = &cli.Command{
	Name:    "recording-failed",
	Aliases: []string{"recorded-failed"},
	Usage:   "recordingFailedCommand で使用",
	Description: `
   録画中のエラー発生時に実行するコマンド．
   recordingFailedCommand で使用する．
`,
	Action: recordingFailedCommandAction,
}

func recordingFailedCommandAction(context *cli.Context) error {
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

	commandConfig := config.Commands.RecordingFailed
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
