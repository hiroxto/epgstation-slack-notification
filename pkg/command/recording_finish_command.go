package command

import (
	"fmt"

	"github.com/hiroxto/epgstation-slack-notification/pkg/app"
	"github.com/hiroxto/epgstation-slack-notification/pkg/config"
	"github.com/hiroxto/epgstation-slack-notification/pkg/env"
	"github.com/hiroxto/epgstation-slack-notification/pkg/service"
	"github.com/urfave/cli/v2"
)

var RecordingFinishCommand = &cli.Command{
	Name:    "recording-finish",
	Aliases: []string{"recorded-end"},
	Usage:   "recordingFinishCommand で使用",
	Description: `
   録画終了時に実行するコマンド．
   recordingFinishCommand で使用する．
`,
	Action: recordingFinishCommandAction,
}

func recordingFinishCommandAction(context *cli.Context) error {
	fileService := service.NewFileService()
	configData, err := fileService.ReadConfigFile(context.String("config"))
	if err != nil {
		return err
	}

	conf, err := config.LoadConfigFromYaml(configData)
	if err != nil {
		return err
	}

	commandConfig := conf.Commands.RecordingFinish
	if !commandConfig.Enable {
		fmt.Printf("%s command is disabled.\n", context.Command.Name)
		return nil
	}
	slackAPIKey := conf.Slack.APIKey
	slackChannel := conf.Slack.Channel
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
		UserName:        commandConfig.UserName,
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
