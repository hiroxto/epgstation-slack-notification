package command

import (
	"os"

	"github.com/urfave/cli/v2"
)

// getConfigFilePath 設定ファイルのパスを取得する
func getConfigFilePath(context *cli.Context) string {
	return context.String("config")
}

// readConfigFile 設定ファイルを読み込む
func readConfigFile(configFilePath string) ([]byte, error) {
	data, err := os.ReadFile(configFilePath)
	if err != nil {
		return nil, err
	}

	return data, err
}
