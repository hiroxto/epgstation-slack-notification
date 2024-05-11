package app

import (
	"strconv"
	"time"

	"github.com/hiroxto/epgstation-slack-notification/pkg/env"
)

// ReserveUseCaseParam ReserveFinishNotificationUseCaseのパラメータ
type ReserveUseCaseParam struct {
	EnableDebug   bool
	SlackAPIKey   string
	SlackChannel  string
	Message       string
	Fields        []Field
	ReserveDetail ReserveDetail
}

// ReserveDetail 予約情報
type ReserveDetail struct {
	ProgramID            string
	ChannelType          string
	ChannelID            string
	ChannelName          string
	HalfWidthChannelName string
	StartAt              string
	EndAt                string
	Duration             uint
	Name                 string
	HalfWidthName        string
	Description          string
	HalfWidthDescription string
	Extended             string
	HalfWidthExtended    string
	StartAtTime          time.Time
	EndAtTime            time.Time
	DurationMin          uint
}

// ReserveDetailFromEnv env.ReserveCommandEnv を ReserveDetail に変換する
func ReserveDetailFromEnv(reserveEnv env.ReserveCommandEnv) ReserveDetail {
	startAt, _ := strconv.ParseInt(reserveEnv.StartAt, 10, 64)
	startAtTime := time.UnixMilli(startAt)
	endAt, _ := strconv.ParseInt(reserveEnv.EndAt, 10, 64)
	endAtTime := time.UnixMilli(endAt)
	durationMs, _ := strconv.ParseUint(reserveEnv.Duration, 10, 64)
	durationInSeconds := durationMs / 1000
	durationMin := durationInSeconds / 60

	return ReserveDetail{
		ProgramID:            reserveEnv.ProgramID,
		ChannelType:          reserveEnv.ChannelType,
		ChannelID:            reserveEnv.ChannelID,
		ChannelName:          reserveEnv.ChannelName,
		HalfWidthChannelName: reserveEnv.HalfWidthChannelName,
		StartAt:              reserveEnv.StartAt,
		EndAt:                reserveEnv.EndAt,
		Duration:             uint(durationMs),
		Name:                 reserveEnv.Name,
		HalfWidthName:        reserveEnv.HalfWidthName,
		Description:          reserveEnv.Description,
		HalfWidthDescription: reserveEnv.HalfWidthDescription,
		Extended:             reserveEnv.Extended,
		HalfWidthExtended:    reserveEnv.HalfWidthExtended,
		StartAtTime:          startAtTime,
		EndAtTime:            endAtTime,
		DurationMin:          uint(durationMin),
	}
}

// ReserveNotificationUseCase 予約関連を通知する
func ReserveNotificationUseCase(param ReserveUseCaseParam) error {
	slackClient, err := createSlackClient(param.SlackAPIKey, param.EnableDebug)
	if err != nil {
		return err
	}

	message, err := formatContent("", param.Message, param.ReserveDetail)
	if err != nil {
		return err
	}

	fields, err := buildCommandFields(param.Fields, param.ReserveDetail)
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
