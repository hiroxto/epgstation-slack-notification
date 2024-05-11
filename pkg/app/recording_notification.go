package app

import (
	"strconv"
	"time"

	"github.com/hiroxto/epgstation-slack-notification/pkg/env"
)

// RecordingUseCaseParam RecordingNotificationUseCaseのパラメータ
type RecordingUseCaseParam struct {
	EnableDebug     bool
	SlackAPIKey     string
	SlackChannel    string
	Message         string
	Fields          []Field
	RecordingDetail RecordingDetail
}

// RecordingDetail 録画情報
type RecordingDetail struct {
	RecordedID           string
	ProgramID            string
	ChannelType          string
	ChannelID            string
	ChannelName          string
	HalfWidthChannelName string
	StartAt              string
	EndAt                string
	Duration             string
	Name                 string
	HalfWidthName        string
	Description          string
	HalfWidthDescription string
	Extended             string
	HalfWidthExtended    string
	RecPath              string
	LogPath              string
	ErrorCnt             int
	DropCnt              int
	ScramblingCount      int
	StartAtTime          time.Time
	EndAtTime            time.Time
	DurationMin          int64
}

// RecordingDetailFromEnv env.RecordingCommandEnv を RecordingDetail に変換する
func RecordingDetailFromEnv(recordingEnv env.RecordingCommandEnv) RecordingDetail {
	startAt, _ := strconv.ParseInt(recordingEnv.StartAt, 10, 64)
	startAtTime := time.UnixMilli(startAt)
	endAt, _ := strconv.ParseInt(recordingEnv.EndAt, 10, 64)
	endAtTime := time.UnixMilli(endAt)
	duration, _ := strconv.ParseInt(recordingEnv.Duration, 10, 64)
	durationInSeconds := duration / 1000
	durationMin := durationInSeconds / 60

	return RecordingDetail{
		RecordedID:           recordingEnv.RecordedID,
		ProgramID:            recordingEnv.ProgramID,
		ChannelType:          recordingEnv.ChannelType,
		ChannelID:            recordingEnv.ChannelID,
		ChannelName:          recordingEnv.ChannelName,
		HalfWidthChannelName: recordingEnv.HalfWidthChannelName,
		StartAt:              recordingEnv.StartAt,
		EndAt:                recordingEnv.EndAt,
		Duration:             recordingEnv.Duration,
		Name:                 recordingEnv.Name,
		HalfWidthName:        recordingEnv.HalfWidthName,
		Description:          recordingEnv.Description,
		HalfWidthDescription: recordingEnv.HalfWidthDescription,
		Extended:             recordingEnv.Extended,
		HalfWidthExtended:    recordingEnv.HalfWidthExtended,
		RecPath:              recordingEnv.RecPath,
		LogPath:              recordingEnv.LogPath,
		ErrorCnt:             recordingEnv.ErrorCnt,
		DropCnt:              recordingEnv.DropCnt,
		ScramblingCount:      recordingEnv.ScramblingCount,
		StartAtTime:          startAtTime,
		EndAtTime:            endAtTime,
		DurationMin:          durationMin,
	}
}

// RecordingNotificationUseCase 予約関連を通知する
func RecordingNotificationUseCase(param RecordingUseCaseParam) error {
	slackClient, err := createSlackClient(param.SlackAPIKey, param.EnableDebug)
	if err != nil {
		return err
	}

	message, err := formatContent("", param.Message, param.RecordingDetail)
	if err != nil {
		return err
	}

	fields, err := buildCommandFields(param.Fields, param.RecordingDetail)
	if err != nil {
		return err
	}

	options := buildMessageOptions(message, fields)
	_, _, err = slackClient.PostMessage(param.SlackChannel, options)

	if err != nil {
		return err
	}

	return nil
}
