package command

import (
	"fmt"

	"github.com/hiroxto/epgstation-slack-notification/pkg/env"
	"github.com/k0kubun/pp/v3"
	"github.com/urfave/cli/v2"
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
		&cli.StringSliceFlag{
			Name:  "only",
			Value: cli.NewStringSlice(dumpEnvCommandValidOnlyValues...),
		},
		&cli.BoolFlag{
			Name:  "color",
			Value: false,
		},
	},
}

var dumpEnvCommandValidOnlyValues = []string{"reserve", "recording", "encoding"}

func dumpEnvCommandAction(context *cli.Context) error {
	onlyValues := context.StringSlice("only")

	// onlyオプションのチェック
	invalidOnlyOptions := []string{}
	for _, onlyValue := range onlyValues {
		found := false
		for _, allowedOnlyValue := range dumpEnvCommandValidOnlyValues {
			if onlyValue == allowedOnlyValue {
				found = true
				break
			}
		}

		if !found {
			invalidOnlyOptions = append(invalidOnlyOptions, onlyValue)
		}
	}
	if len(invalidOnlyOptions) != 0 {
		return fmt.Errorf("invalid only options : %v", invalidOnlyOptions)
	}

	pp.Default.SetColoringEnabled(context.Bool("color"))
	for _, onlyValue := range onlyValues {
		switch onlyValue {
		case "reserve":
			reserveCommandEnv, err := env.LoadReserveCommandEnv()
			if err != nil {
				return err
			}
			pp.Println(reserveCommandEnv)
		case "recording":
			recordingCommandEnv, err := env.LoadRecordingCommandEnv()
			if err != nil {
				return err
			}
			pp.Println(recordingCommandEnv)
		case "encoding":
			encodingCommandEnv, err := env.LoadEncodingCommandEnv()
			if err != nil {
				return err
			}
			pp.Println(encodingCommandEnv)
		default:
			return fmt.Errorf("unknown value:%v", onlyValue)
		}
	}

	return nil
}
