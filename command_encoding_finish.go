package main

import (
	"github.com/urfave/cli/v2"
)

var commandEncodingFinish = &cli.Command{
	Name:  "encoding-finish",
	Usage: "encodingFinishCommand で使用",
	Description: `
   エンコード終了時に実行するコマンド。
   encodingFinishCommand で使用する。
`,
	Action: commandEncodingFinishAction,
}

func commandEncodingFinishAction(context *cli.Context) error {

	return nil
}
