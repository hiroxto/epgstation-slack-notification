package command

import (
	"fmt"

	"github.com/hiroxto/epgstation-slack-notification/pkg/env"
	"github.com/urfave/cli/v2"
)

// DumpEnvsCommand dump-envsコマンド
var DumpEnvsCommand = &cli.Command{
	Name:  "dump-envs",
	Usage: "環境変数を出力するデバッグ用コマンド",
	Description: `
   環境変数を出力するデバッグ用コマンド
`,
	Action: dumpEnvsCommandAction,
	Flags: []cli.Flag{
		&cli.StringSliceFlag{
			Name:  "only",
			Value: cli.NewStringSlice(dumpEnvsCommandValidOnlyValues...),
		},
	},
}

var dumpEnvsCommandValidOnlyValues = []string{"reserve", "recording", "encoding"}

func dumpEnvsCommandAction(context *cli.Context) error {
	onlyValues := context.StringSlice("only")

	// onlyオプションのチェック
	invalidOnlyOptions := []string{}
	for _, onlyValue := range onlyValues {
		found := false
		for _, allowedOnlyValue := range dumpEnvsCommandValidOnlyValues {
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

	for _, onlyValue := range onlyValues {
		switch onlyValue {
		case "reserve":
			recordingCommandEnv, err := env.LoadRecordingCommandEnv()
			if err != nil {
				return err
			}
			fmt.Printf("%#v\n", recordingCommandEnv)
		case "recording":
			recordingCommandEnv, err := env.LoadRecordingCommandEnv()
			if err != nil {
				return err
			}
			fmt.Printf("%#v\n", recordingCommandEnv)
		case "encoding":
			encodingCommandEnv, err := env.LoadEncodingCommandEnv()
			if err != nil {
				return err
			}
			fmt.Printf("%#v\n", encodingCommandEnv)
		default:
			return fmt.Errorf("unknown value:%v", onlyValue)
		}
	}

	return nil
}
