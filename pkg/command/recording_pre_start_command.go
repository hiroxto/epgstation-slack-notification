package command

import (
	"context"
	"fmt"

	"github.com/hiroxto/epgstation-slack-notification/pkg/app"
	"github.com/hiroxto/epgstation-slack-notification/pkg/config"
	"github.com/hiroxto/epgstation-slack-notification/pkg/env"
	"github.com/hiroxto/epgstation-slack-notification/pkg/service"
	"github.com/urfave/cli/v3"
)

var RecordingPreStartCommand = &cli.Command{
	Name:    "recording-pre-start",
	Aliases: []string{"recorded-pre-start"},
	Usage:   "recordingPreStartCommand で使用",
	Description: `
   録画準備の開始時に実行されるコマンド．
   recordingPreStartCommand で使用する．
`,
	Action: recordingPreStartCommandAction,
}

func recordingPreStartCommandAction(_ context.Context, cmd *cli.Command) error {
	fileService := service.NewFileService()
	configData, err := fileService.ReadConfigFile(cmd.String("config"))
	if err != nil {
		return err
	}

	conf, err := config.LoadConfigFromYaml(configData)
	if err != nil {
		return err
	}

	commandConfig := conf.Commands.RecordingPreStart
	if !commandConfig.Enable {
		fmt.Printf("%s command is disabled.\n", cmd.Name)
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
		EnableDebug:   cmd.Bool("debug"),
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
