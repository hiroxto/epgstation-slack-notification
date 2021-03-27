package main

import (
	"github.com/urfave/cli/v2"
)

var commandReserveDeleted = &cli.Command{
	Name:  "reserve-deleted",
	Usage: "reservedeletedCommand で使用",
	Description: `
   録画予約の削除時に実行されるコマンド．
   reservedeletedCommand で使用する．
`,
	Action: commandReserveDeletedAction,
}

func commandReserveDeletedAction(context *cli.Context) error {
	config, err := loadConfigFile()
	if err != nil {
		return err
	}

	env, err := loadCommandEnv()
	if err != nil {
		return err
	}

	err = startCommandNotification(context, env, config, config.Commands.ReserveDeletedCommand)
	if err != nil {
		return err
	}

	return nil
}
