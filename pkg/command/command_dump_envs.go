package command

import (
	"fmt"

	"github.com/hiroxto/epgstation-slack-notification/pkg/env"
	"github.com/urfave/cli/v2"
)

// CommandDumpEnvs dump-envsコマンド
var CommandDumpEnvs = &cli.Command{
	Name:  "dump-envs",
	Usage: "環境変数を出力するデバッグ用コマンド",
	Description: `
   環境変数を出力するデバッグ用コマンド
`,
	Action: commandDumpEnvsAction,
}

func commandDumpEnvsAction(context *cli.Context) error {
	reserveCommandEnv, err := env.LoadReserveCommandEnv()
	if err != nil {
		return err
	}
	fmt.Printf("%#v\n", reserveCommandEnv)

	recordingCommandEnv, err := env.LoadRecordingCommandEnv()
	if err != nil {
		return err
	}
	fmt.Printf("%#v\n", recordingCommandEnv)

	encodingCommandEnv, err := env.LoadEncodingCommandEnv()
	if err != nil {
		return err
	}
	fmt.Printf("%#v\n", encodingCommandEnv)

	return nil
}
