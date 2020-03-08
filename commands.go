package main

import (
	"fmt"
	"github.com/urfave/cli/v2"
)

var commands = []*cli.Command{
	commandReservationAdded,
	commandRecordedPreStart,
	commandRecordedPrepRecFailed,

	commandRecordedStart,
	commandRecordedEnd,
	commandRecordedFailed,
}

func displayCommandIsDisableMessage(context *cli.Context) {
	fmt.Printf("%s command is disabled.\n", context.Command.Name)
}
