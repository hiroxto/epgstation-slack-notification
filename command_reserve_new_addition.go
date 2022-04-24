package main

import (
	"github.com/urfave/cli/v2"
)

var commandReserveNewAddition = &cli.Command{
	Name:    "reserve-new-addition",
	Aliases: []string{"reservation-added"},
	Usage:   "reserveNewAddtionCommand で使用",
	Description: `
   録画予約の新規追加時に実行されるコマンド．
   reserveNewAddtionCommand で使用する．
`,
	Action: commandReserveNewAdditionAction,
}

func commandReserveNewAdditionAction(context *cli.Context) error {
	config, err := loadConfigFile()
	if err != nil {
		return err
	}

	var env RecordingCommandEnv
	if err := loadCommandEnv(&env); err != nil {
		return err
	}

	err = startCommandNotification(context, env, config, config.Commands.ReserveNewAddition)
	if err != nil {
		return err
	}

	return nil
}
