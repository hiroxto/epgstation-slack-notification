package service

import (
	"os"
)

type FileService struct{}

func NewFileService() *FileService {
	return &FileService{}
}

// ReadConfigFile 設定ファイルを読み込む
func (fs *FileService) ReadConfigFile(configFilePath string) ([]byte, error) {
	data, err := os.ReadFile(configFilePath)
	if err != nil {
		return nil, err
	}

	return data, nil
}
