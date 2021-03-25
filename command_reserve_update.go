package main

import (
	"github.com/urfave/cli/v2"
)

var commandReserveUpdate = &cli.Command{
	Name:  "reserve-update",
	Usage: "reserveUpdateCommand で使用",
	Description: `
   録画情報の更新時に実行されるコマンド．
   reserveUpdateCommand で使用する．
`,
	Action: commandReserveUpdateAction,
}

func commandReserveUpdateAction(context *cli.Context) error {
	config, err := loadConfigFile()
	if err != nil {
		return err
	}

	env, err := loadCommandEnv()
	if err != nil {
		return err
	}

	err = startCommandNotification(context, env, config, config.Commands.ReserveUpdateCommand)
	if err != nil {
		return err
	}

	return nil
}
