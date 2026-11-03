package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/hiroxto/epgstation-slack-notification/pkg/command"
	"github.com/urfave/cli/v3"
)

// DefaultConfigFileName デフォルトの設定ファイル名
const DefaultConfigFileName = "epgstation-slack-config.yml"

var (
	version = "dev"
	commit  = "unknown"
	date    = "unknown"
)

func main() {
	app, err := newCliApp()
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	err = app.Run(context.Background(), os.Args)
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

func newCliApp() (*cli.Command, error) {
	// 同一ディレクトリにあるファイルを指定
	exeFilePath, err := os.Executable()
	if err != nil {
		return nil, err
	}
	defaultConfigFilePath := filepath.Join(filepath.Dir(exeFilePath), DefaultConfigFileName)

	app := &cli.Command{}
	app.Name = "epgstation-slack-notification"
	app.Usage = "EPGStationの通知をSlackに送るコマンドラインツール"
	app.Version = fmt.Sprintf("%s (commit: %s, date: %s)", version, commit, date)
	app.Copyright = "(c) 2020 Hiroto Kitazawa"
	app.Commands = commands
	app.Flags = []cli.Flag{
		&cli.BoolFlag{Name: "debug", Usage: "デバッグモードを有効化"},
		&cli.StringFlag{
			Name: "config",
			Aliases: []string{
				"c",
			},
			Usage: "設定ファイルのパスを指定。",
			Value: defaultConfigFilePath,
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

	command.DumpConfigCommand,
	command.DumpDetailCommand,
	command.DumpEnvCommand,
}
