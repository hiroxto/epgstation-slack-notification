package main

import (
	"github.com/urfave/cli/v2"
)

var commandReservationAdded = &cli.Command{
	Name:  "reservation-added",
	Usage: "reservationAddedCommand で使用",
	Description: `
   録画予約の新規追加時に実行されるコマンド.
   reservationAddedCommand で使用する.
`,
	Action: commandReservationAddedAction,
}

func commandReservationAddedAction(context *cli.Context) error {
	config, err := loadConfigFile()
	if err != nil {
		return err
	}

	env, err := loadPreCommandEnvs()
	if err != nil {
		return err
	}

	err = startPreCommandNotification(context, env, config, config.Commands.ReservationAdded)
	if err != nil {
		return err
	}

	return nil
}
