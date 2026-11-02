package command

import (
	"context"
	"github.com/hiroxto/epgstation-slack-notification/pkg/app"
	"github.com/hiroxto/epgstation-slack-notification/pkg/env"
	"github.com/k0kubun/pp/v3"
	"github.com/urfave/cli/v3"
)

// DumpDetailCommand dump:detail コマンド
var DumpDetailCommand = &cli.Command{
	Name:  "dump:detail",
	Usage: "template へ渡されるデータを出力するデバッグ用コマンド",
	Description: `
   template へ渡されるデータを出力するデバッグ用コマンド
`,
	Action: dumpDetailCommandAction,
	Flags: []cli.Flag{
		&cli.BoolFlag{
			Name:  "color",
			Value: false,
		},
	},
}

func dumpDetailCommandAction(_ context.Context, cmd *cli.Command) error {
	pp.Default.SetColoringEnabled(cmd.Bool("color"))

	// reserve
	reserveCommandEnv, err := env.LoadReserveCommandEnv()
	if err != nil {
		return err
	}
	reserveDetail := app.ReserveDetailFromEnv(reserveCommandEnv)
	pp.Println(reserveDetail)

	// recording
	recordingCommandEnv, err := env.LoadRecordingCommandEnv()
	if err != nil {
		return err
	}
	recordingDetail := app.RecordingDetailFromEnv(recordingCommandEnv)
	pp.Println(recordingDetail)

	// encoding
	encodingCommandEnv, err := env.LoadEncodingCommandEnv()
	if err != nil {
		return err
	}
	encodingDetail := app.EncodingDetailFromEnv(encodingCommandEnv)
	pp.Println(encodingDetail)

	return nil
}
