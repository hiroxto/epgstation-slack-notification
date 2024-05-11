package app

import (
	"reflect"
	"testing"
	"time"

	"github.com/hiroxto/epgstation-slack-notification/pkg/env"
)

func Test_RecordingDetailFromEnv(t *testing.T) {
	env := env.RecordingCommandEnv{
		ProgramID:            1,
		ChannelType:          "CHANNELTYPE",
		ChannelID:            2,
		ChannelName:          "CHANNELNAME",
		HalfWidthChannelName: "HALF_WIDTH_CHANNELNAME",
		StartAt:              1715353200000, // Unix Millis = 2024-05-11T00:00:00+09:00
		EndAt:                1715355000000, // Unix Millis = 2024-05-11T00:30:00+09:00
		Duration:             1800000,       // ms = 30 min
		Name:                 "NAME",
		HalfWidthName:        "HALF_WIDTH_NAME",
		Description:          "DESCRIPTION",
		HalfWidthDescription: "HALF_WIDTH_DESCRIPTION",
		Extended:             "EXTENDED",
		HalfWidthExtended:    "HALF_WIDTH_EXTENDED",
		RecPath:              "RECPATH",
		LogPath:              "LOGPATH",
		ErrorCnt:             3,
		DropCnt:              4,
		ScramblingCount:      5,
	}

	actual := RecordingDetailFromEnv(env)

	expected := RecordingDetail{
		ProgramID:            1,
		ChannelType:          "CHANNELTYPE",
		ChannelID:            2,
		ChannelName:          "CHANNELNAME",
		HalfWidthChannelName: "HALF_WIDTH_CHANNELNAME",
		StartAt:              1715353200000,
		EndAt:                1715355000000,
		Duration:             1800000,
		Name:                 "NAME",
		HalfWidthName:        "HALF_WIDTH_NAME",
		Description:          "DESCRIPTION",
		HalfWidthDescription: "HALF_WIDTH_DESCRIPTION",
		Extended:             "EXTENDED",
		HalfWidthExtended:    "HALF_WIDTH_EXTENDED",
		RecPath:              "RECPATH",
		LogPath:              "LOGPATH",
		ErrorCnt:             3,
		DropCnt:              4,
		ScramblingCount:      5,
		StartAtTime:          time.UnixMilli(1715353200000),
		EndAtTime:            time.UnixMilli(1715355000000),
		DurationMin:          30,
	}

	if !reflect.DeepEqual(expected, actual) {
		t.Fatalf("expected = %v, actual = %v", expected, actual)
	}

}
