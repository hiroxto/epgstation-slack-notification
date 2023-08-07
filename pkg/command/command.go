package command

import (
	"os"
	"path/filepath"

	"github.com/hiroxto/epgstation-slack-notification/pkg/app"
	"github.com/hiroxto/epgstation-slack-notification/pkg/env"
)

// DefaultConfigFileName デフォルトの設定ファイル名
const DefaultConfigFileName = "epgstation-slack-config.yml"

var executable = os.Executable

func getConfigFilePath() (string, error) {
	exeFilePath, err := executable()
	if err != nil {
		return exeFilePath, err
	}

	return filepath.Join(filepath.Dir(exeFilePath), DefaultConfigFileName), nil
}

// encodingEnvToEncodingFinishDetail env.EncodingCommandEnvをapp.EncodingFinishDetailに変換する
func encodingEnvToEncodingFinishDetail(encodingEnv env.EncodingCommandEnv) app.EncodingFinishDetail {
	return app.EncodingFinishDetail{
		RecordedID:  encodingEnv.RecordedID,
		VideoFileID: encodingEnv.VideoFileID,
		OutputPath:  encodingEnv.OutputPath,
		Mode:        encodingEnv.Mode,
		ChannelID:   encodingEnv.ChannelID,
		ChannelName: encodingEnv.HalfWidthChannelName,
		Name:        encodingEnv.HalfWidthName,
		Description: encodingEnv.HalfWidthDescription,
		Extended:    encodingEnv.HalfWidthExtended,
		Original: app.EncodingFinishOriginal{
			RecordedID:           encodingEnv.RecordedID,
			VideoFileID:          encodingEnv.VideoFileID,
			OutputPath:           encodingEnv.OutputPath,
			Mode:                 encodingEnv.Mode,
			ChannelID:            encodingEnv.ChannelID,
			ChannelName:          encodingEnv.ChannelName,
			HalfWidthChannelName: encodingEnv.HalfWidthChannelName,
			Name:                 encodingEnv.Name,
			HalfWidthName:        encodingEnv.HalfWidthName,
			Description:          encodingEnv.Description,
			HalfWidthDescription: encodingEnv.HalfWidthDescription,
			Extended:             encodingEnv.Extended,
			HalfWidthExtended:    encodingEnv.HalfWidthExtended,
		},
	}
}
