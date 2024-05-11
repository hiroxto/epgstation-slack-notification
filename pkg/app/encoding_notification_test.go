package app

import (
	"reflect"
	"testing"

	"github.com/hiroxto/epgstation-slack-notification/pkg/env"
)

func Test_EncodingDetailFromEnv(t *testing.T) {
	env := env.EncodingCommandEnv{
		RecordedID:           "1",
		VideoFileID:          "2",
		OutputPath:           "OUTPUTPATH",
		Mode:                 "MODE",
		ChannelID:            "3",
		ChannelName:          "CHANNELNAME",
		HalfWidthChannelName: "HALF_WIDTH_CHANNELNAME",
		Name:                 "NAME",
		HalfWidthName:        "HALF_WIDTH_NAME",
		Description:          "DESCRIPTION",
		HalfWidthDescription: "HALF_WIDTH_DESCRIPTION",
		Extended:             "EXTENDED",
		HalfWidthExtended:    "HALF_WIDTH_EXTENDED",
	}

	actual := EncodingDetailFromEnv(env)

	expected := EncodingDetail{
		RecordedID:           "1",
		VideoFileID:          "2",
		OutputPath:           "OUTPUTPATH",
		Mode:                 "MODE",
		ChannelID:            "3",
		ChannelName:          "CHANNELNAME",
		HalfWidthChannelName: "HALF_WIDTH_CHANNELNAME",
		Name:                 "NAME",
		HalfWidthName:        "HALF_WIDTH_NAME",
		Description:          "DESCRIPTION",
		HalfWidthDescription: "HALF_WIDTH_DESCRIPTION",
		Extended:             "EXTENDED",
		HalfWidthExtended:    "HALF_WIDTH_EXTENDED",
	}

	if !reflect.DeepEqual(expected, actual) {
		t.Fatalf("expected = %v, actual = %v", expected, actual)
	}
}
