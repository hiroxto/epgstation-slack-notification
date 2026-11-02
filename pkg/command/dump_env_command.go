package command

import (
	"context"
	"github.com/hiroxto/epgstation-slack-notification/pkg/env"
	"github.com/k0kubun/pp/v3"
	"github.com/urfave/cli/v3"
)

// DumpEnvCommand dump:env コマンド
var DumpEnvCommand = &cli.Command{
	Name:  "dump:env",
	Usage: "環境変数を出力するデバッグ用コマンド",
	Description: `
   環境変数を出力するデバッグ用コマンド
`,
	Action: dumpEnvCommandAction,
	Flags: []cli.Flag{
		&cli.BoolFlag{
			Name:  "color",
			Value: false,
		},
	},
}

func dumpEnvCommandAction(_ context.Context, cmd *cli.Command) error {
	pp.Default.SetColoringEnabled(cmd.Bool("color"))

	// reserve
	reserveCommandEnv, err := env.LoadReserveCommandEnv()
	if err != nil {
		return err
	}
	pp.Println(reserveCommandEnv)

	// recording
	recordingCommandEnv, err := env.LoadRecordingCommandEnv()
	if err != nil {
		return err
	}
	pp.Println(recordingCommandEnv)

	// encoding
	encodingCommandEnv, err := env.LoadEncodingCommandEnv()
	if err != nil {
		return err
	}
	pp.Println(encodingCommandEnv)

	return nil
}
