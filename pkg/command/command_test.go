package command

import (
	"reflect"
	"testing"

	"github.com/hiroxto/epgstation-slack-notification/pkg/app"
	"github.com/hiroxto/epgstation-slack-notification/pkg/env"
)

func Test_getConfigFilePath_設定ファイルのパスを読み込める(t *testing.T) {
	old := executable
	defer func() { executable = old }()
	executable = func() (string, error) {
		return "/path/to/executable", nil
	}

	actual, err := getConfigFilePath()
	if err != nil {
		t.Error(err)
	}

	expected := "/path/to/epgstation-slack-config.yml"

	if !reflect.DeepEqual(expected, actual) {
		t.Fatalf("expected = %v, actual = %v", expected, actual)
	}
}

func Test_encodingEnvToEncodingFinishDetail(t *testing.T) {
	env := env.EncodingCommandEnv{
		RecordedID:           "RECORDEDID",
		VideoFileID:          "VIDEOFILEID",
		OutputPath:           "OUTPUTPATH",
		Mode:                 "MODE",
		ChannelID:            "CHANNELID",
		ChannelName:          "CHANNELNAME",
		HalfWidthChannelName: "HALF_WIDTH_CHANNELNAME",
		Name:                 "NAME",
		HalfWidthName:        "HALF_WIDTH_NAME",
		Description:          "DESCRIPTION",
		HalfWidthDescription: "HALF_WIDTH_DESCRIPTION",
		Extended:             "EXTENDED",
		HalfWidthExtended:    "HALF_WIDTH_EXTENDED",
	}

	actual := encodingEnvToEncodingFinishDetail(env)

	expected := app.EncodingFinishDetail{
		RecordedID:  "RECORDEDID",
		VideoFileID: "VIDEOFILEID",
		OutputPath:  "OUTPUTPATH",
		Mode:        "MODE",
		ChannelID:   "CHANNELID",
		ChannelName: "HALF_WIDTH_CHANNELNAME",
		Name:        "HALF_WIDTH_NAME",
		Description: "HALF_WIDTH_DESCRIPTION",
		Extended:    "HALF_WIDTH_EXTENDED",
		Original: app.EncodingFinishOriginal{
			RecordedID:           "RECORDEDID",
			VideoFileID:          "VIDEOFILEID",
			OutputPath:           "OUTPUTPATH",
			Mode:                 "MODE",
			ChannelID:            "CHANNELID",
			ChannelName:          "CHANNELNAME",
			HalfWidthChannelName: "HALF_WIDTH_CHANNELNAME",
			Name:                 "NAME",
			HalfWidthName:        "HALF_WIDTH_NAME",
			Description:          "DESCRIPTION",
			HalfWidthDescription: "HALF_WIDTH_DESCRIPTION",
			Extended:             "EXTENDED",
			HalfWidthExtended:    "HALF_WIDTH_EXTENDED",
		},
	}

	if !reflect.DeepEqual(expected, actual) {
		t.Fatalf("expected = %v, actual = %v", expected, actual)
	}
}
