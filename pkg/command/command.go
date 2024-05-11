package command

import (
	"os"
	"path/filepath"
)

// DefaultConfigFileName デフォルトの設定ファイル名
const DefaultConfigFileName = "epgstation-slack-config.yml"

var executable = os.Executable

// getConfigFilePath 設定ファイルのパスを取得する
// バイナリと同じディレクトリにあるDefaultConfigFileNameのファイル名を探す
func getConfigFilePath() (string, error) {
	exeFilePath, err := executable()
	if err != nil {
		return exeFilePath, err
	}

	return filepath.Join(filepath.Dir(exeFilePath), DefaultConfigFileName), nil
}
