package command

import (
	"fmt"

	"github.com/hiroxto/epgstation-slack-notification/pkg/app"
	"github.com/hiroxto/epgstation-slack-notification/pkg/config"
	"github.com/hiroxto/epgstation-slack-notification/pkg/env"
	"github.com/hiroxto/epgstation-slack-notification/pkg/service"
	"github.com/urfave/cli/v2"
)

var RecordingPrepRecFailedCommand = &cli.Command{
	Name:    "recording-prep-rec-failed",
	Aliases: []string{"recorded-prep-rec-failed"},
	Usage:   "recordingPrepRecFailedCommand で使用",
	Description: `
   録画準備の失敗時に実行されるコマンド．
   recordingPrepRecFailedCommand で使用する．`,
	Action: recordingPrepRecFailedCommandAction,
}

func recordingPrepRecFailedCommandAction(context *cli.Context) error {
	fileService := service.NewFileService()
	configData, err := fileService.ReadConfigFile(context.String("config"))
	if err != nil {
		return err
	}

	conf, err := config.LoadConfigFromYaml(configData)
	if err != nil {
		return err
	}

	commandConfig := conf.Commands.RecordingPrepRecFailed
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
