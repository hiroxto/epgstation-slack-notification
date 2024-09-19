package app

import (
	"github.com/hiroxto/epgstation-slack-notification/pkg/env"
	"github.com/hiroxto/epgstation-slack-notification/pkg/service"
)

// EncodingUseCaseParam EncodingUseCaseのパラメータ
type EncodingUseCaseParam struct {
	EnableDebug    bool
	SlackAPIKey    string
	SlackChannel   string
	UserName       string
	Message        string
	Fields         []Field
	EncodingDetail EncodingDetail
}

// EncodingDetail エンコーディング情報
type EncodingDetail struct {
	RecordedID           string
	VideoFileID          string
	OutputPath           string
	Mode                 string
	ChannelID            string
	ChannelName          string
	HalfWidthChannelName string
	Name                 string
	HalfWidthName        string
	Description          string
	HalfWidthDescription string
	Extended             string
	HalfWidthExtended    string
}

// EncodingDetailFromEnv env.EncodingCommandEnv を EncodingDetail に変換する
func EncodingDetailFromEnv(encodingEnv env.EncodingCommandEnv) EncodingDetail {
	return EncodingDetail{
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
	}
}

// EncodingNotificationUseCase エンコーディング関連を通知する
func EncodingNotificationUseCase(param EncodingUseCaseParam) error {
	client := service.NewSlackService(param.SlackAPIKey, param.EnableDebug)

	err := client.PostMessageWithFields(param.SlackChannel, param.Message, convertFields(param.Fields), param.EncodingDetail, param.UserName)
	if err != nil {
		return err
	}

	return nil
}
