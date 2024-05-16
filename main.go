package main

import (
	"log"
	"os"
	"path/filepath"

	"github.com/hiroxto/epgstation-slack-notification/pkg/command"
	"github.com/urfave/cli/v2"
)

// DefaultConfigFileName デフォルトの設定ファイル名
const DefaultConfigFileName = "epgstation-slack-config.yml"

func main() {
	app, err := newCliApp()
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	err = app.Run(os.Args)
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

func newCliApp() (*cli.App, error) {
	// 同一ディレクトリにあるファイルを指定
	exeFilePath, err := os.Executable()
	if err != nil {
		return nil, err
	}
	defaultConfigFilePath := filepath.Join(filepath.Dir(exeFilePath), DefaultConfigFileName)

	app := cli.NewApp()
	app.Name = "epgstation-slack-notification"
	app.Usage = "EPGStationの通知をSlackに送るコマンドラインツール"
	app.Version = "2.2.0"
	app.Copyright = "(c) 2020 Hiroto Kitazawa"
	app.Commands = commands
	app.Flags = []cli.Flag{
		&cli.BoolFlag{Name: "debug", Usage: "デバッグモードを有効化"},
		&cli.StringFlag{
			Name: "config",
			Aliases: []string{
				"c",
			},
			Usage:       "設定ファイルのパスを指定。",
			DefaultText: defaultConfigFilePath,
		},
	}

	return app, nil
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
