package command

import (
	"github.com/urfave/cli/v2"
)

// getConfigFilePath 設定ファイルのパスを取得する
func getConfigFilePath(context *cli.Context) string {
	return context.String("config")
}
