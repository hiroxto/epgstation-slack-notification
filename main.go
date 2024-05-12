package main

import (
	"log"
	"os"

	"github.com/hiroxto/epgstation-slack-notification/pkg/command"
	"github.com/urfave/cli/v2"
)

func main() {
	app := newCliApp()
	err := app.Run(os.Args)

	if err != nil {
		log.Println(err)

		exitCode := 1
		excoder, ok := err.(cli.ExitCoder)
		if ok {
			exitCode = excoder.ExitCode()
		}
		os.Exit(exitCode)
	}
}

func newCliApp() *cli.App {
	app := cli.NewApp()
	app.Name = "epgstation-slack-notification"
	app.Usage = "EPGStationの通知をSlackに送るコマンドラインツール"
	app.Version = "2.2.0"
	app.Copyright = "(c) 2020 Hiroto Kitazawa"
	app.Commands = commands
	app.Flags = []cli.Flag{
		&cli.BoolFlag{Name: "debug", Usage: "デバッグモードを有効化"},
	}

	return app
}

var commands = []*cli.Command{
	command.ReserveNewAdditionCommand,
	command.ReserveUpdateCommand,
	command.ReserveDeletedCommand,
	command.RecordingPreStartCommand,
	command.RecordingPrepRecFailedCommand,

	command.RecordingStartCommand,
	command.RecordingFinishCommand,
	command.RecordingFailedCommand,

	command.EncodingFinishCommand,

	command.DumpEnvsCommand,
}
