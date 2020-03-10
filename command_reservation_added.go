package main

import (
	"github.com/urfave/cli/v2"
	"log"
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
		log.Fatal(err)
	}

	env, err := loadPreCommandEnvs()
	if err != nil {
		log.Fatal(err)
	}

	err = startPreCommandNotification(context, env, config, config.Commands.ReservationAdded)

	if err != nil {
		log.Fatal(err)
	}

	return nil
}
