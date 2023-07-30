package command

import (
	"os"
	"path/filepath"
)

const DefaultConfigFileName = "epgstation-slack-config.yml"

var executable = os.Executable

func getConfigFilePath() (string, error) {
	exeFilePath, err := executable()
	if err != nil {
		return exeFilePath, err
	}

	return filepath.Join(filepath.Dir(exeFilePath), DefaultConfigFileName), nil
}
