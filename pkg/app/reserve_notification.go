package app

import (
	"strconv"
	"time"

	"github.com/hiroxto/epgstation-slack-notification/pkg/env"
	"github.com/hiroxto/epgstation-slack-notification/pkg/service"
)

// ReserveUseCaseParam ReserveNotificationUseCase のパラメータ
type ReserveUseCaseParam struct {
	EnableDebug   bool
	SlackAPIKey   string
	SlackChannel  string
	UserName      string
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
	Duration             string
	Name                 string
	HalfWidthName        string
	Description          string
	HalfWidthDescription string
	Extended             string
	HalfWidthExtended    string
	StartAtTime          time.Time
	EndAtTime            time.Time
	DurationMin          int64
}

// ReserveDetailFromEnv env.ReserveCommandEnv を ReserveDetail に変換する
func ReserveDetailFromEnv(reserveEnv env.ReserveCommandEnv) ReserveDetail {
	startAt, _ := strconv.ParseInt(reserveEnv.StartAt, 10, 64)
	startAtTime := time.UnixMilli(startAt)
	endAt, _ := strconv.ParseInt(reserveEnv.EndAt, 10, 64)
	endAtTime := time.UnixMilli(endAt)
	duration, _ := strconv.ParseInt(reserveEnv.Duration, 10, 64)
	durationInSeconds := duration / 1000
	durationMin := durationInSeconds / 60

	return ReserveDetail{
		ProgramID:            reserveEnv.ProgramID,
		ChannelType:          reserveEnv.ChannelType,
		ChannelID:            reserveEnv.ChannelID,
		ChannelName:          reserveEnv.ChannelName,
		HalfWidthChannelName: reserveEnv.HalfWidthChannelName,
		StartAt:              reserveEnv.StartAt,
		EndAt:                reserveEnv.EndAt,
		Duration:             reserveEnv.Duration,
		Name:                 reserveEnv.Name,
		HalfWidthName:        reserveEnv.HalfWidthName,
		Description:          reserveEnv.Description,
		HalfWidthDescription: reserveEnv.HalfWidthDescription,
		Extended:             reserveEnv.Extended,
		HalfWidthExtended:    reserveEnv.HalfWidthExtended,
		StartAtTime:          startAtTime,
		EndAtTime:            endAtTime,
		DurationMin:          durationMin,
	}
}

// ReserveNotificationUseCase 予約関連を通知する
func ReserveNotificationUseCase(param ReserveUseCaseParam) error {
	client := service.NewSlackService(param.SlackAPIKey, param.EnableDebug)

	err := client.PostMessageWithFields(param.SlackChannel, param.Message, convertFields(param.Fields), param.ReserveDetail, param.UserName)
	if err != nil {
		return err
	}

	return nil
}
