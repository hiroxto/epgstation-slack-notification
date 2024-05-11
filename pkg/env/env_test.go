package env

import (
	"os"
	"reflect"
	"testing"
)

func Test_LoadReserveCommandEnv_デフォルトパラメータで読み込める(t *testing.T) {
	os.Clearenv()

	actual, err := LoadReserveCommandEnv()
	if err != nil {
		t.Error(err)
	}

	expected := ReserveCommandEnv{
		ProgramID:            "",
		ChannelType:          "None",
		ChannelID:            "",
		ChannelName:          "None",
		HalfWidthChannelName: "None",
		StartAt:              "",
		EndAt:                "",
		Duration:             "",
		Name:                 "None",
		HalfWidthName:        "None",
		Description:          "None",
		HalfWidthDescription: "None",
		Extended:             "None",
		HalfWidthExtended:    "None",
	}

	if !reflect.DeepEqual(expected, actual) {
		t.Fatalf("expected = %v, actual = %v", expected, actual)
	}
}

func Test_LoadReserveCommandEnv_フルパラメータで読み込める(t *testing.T) {
	os.Clearenv()
	os.Setenv("PROGRAMID", "1")
	os.Setenv("CHANNELTYPE", "CHANNELTYPE")
	os.Setenv("CHANNELID", "2")
	os.Setenv("CHANNELNAME", "CHANNELNAME")
	os.Setenv("HALF_WIDTH_CHANNELNAME", "HALF_WIDTH_CHANNELNAME")
	os.Setenv("STARTAT", "1715353200000")
	os.Setenv("ENDAT", "1715355000000")
	os.Setenv("DURATION", "1800000")
	os.Setenv("NAME", "NAME")
	os.Setenv("HALF_WIDTH_NAME", "HALF_WIDTH_NAME")
	os.Setenv("DESCRIPTION", "DESCRIPTION")
	os.Setenv("HALF_WIDTH_DESCRIPTION", "HALF_WIDTH_DESCRIPTION")
	os.Setenv("EXTENDED", "EXTENDED")
	os.Setenv("HALF_WIDTH_EXTENDED", "HALF_WIDTH_EXTENDED")

	actual, err := LoadReserveCommandEnv()
	if err != nil {
		t.Error(err)
	}

	expected := ReserveCommandEnv{
		ProgramID:            "1",
		ChannelType:          "CHANNELTYPE",
		ChannelID:            "2",
		ChannelName:          "CHANNELNAME",
		HalfWidthChannelName: "HALF_WIDTH_CHANNELNAME",
		StartAt:              "1715353200000",
		EndAt:                "1715355000000",
		Duration:             "1800000",
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

func Test_LoadRecordingCommandEnv_デフォルトパラメータで読み込める(t *testing.T) {
	os.Clearenv()

	actual, err := LoadRecordingCommandEnv()
	if err != nil {
		t.Error(err)
	}

	expected := RecordingCommandEnv{
		RecordedID:           "",
		ProgramID:            "",
		ChannelType:          "None",
		ChannelID:            "",
		ChannelName:          "None",
		HalfWidthChannelName: "None",
		StartAt:              "",
		EndAt:                "",
		Duration:             "",
		Name:                 "None",
		HalfWidthName:        "None",
		Description:          "None",
		HalfWidthDescription: "None",
		Extended:             "None",
		HalfWidthExtended:    "None",
		RecPath:              "None",
		LogPath:              "None",
		ErrorCnt:             "",
		DropCnt:              "",
		ScramblingCount:      "",
	}

	if !reflect.DeepEqual(expected, actual) {
		t.Fatalf("expected = %v, actual = %v", expected, actual)
	}
}

func Test_LoadRecordingCommandEnv_フルパラメータで読み込める(t *testing.T) {
	os.Clearenv()
	os.Setenv("RECORDEDID", "1")
	os.Setenv("PROGRAMID", "2")
	os.Setenv("CHANNELTYPE", "CHANNELTYPE")
	os.Setenv("CHANNELID", "3")
	os.Setenv("CHANNELNAME", "CHANNELNAME")
	os.Setenv("HALF_WIDTH_CHANNELNAME", "HALF_WIDTH_CHANNELNAME")
	os.Setenv("STARTAT", "1715353200000")
	os.Setenv("ENDAT", "1715355000000")
	os.Setenv("DURATION", "1800000")
	os.Setenv("NAME", "NAME")
	os.Setenv("HALF_WIDTH_NAME", "HALF_WIDTH_NAME")
	os.Setenv("DESCRIPTION", "DESCRIPTION")
	os.Setenv("HALF_WIDTH_DESCRIPTION", "HALF_WIDTH_DESCRIPTION")
	os.Setenv("EXTENDED", "EXTENDED")
	os.Setenv("HALF_WIDTH_EXTENDED", "HALF_WIDTH_EXTENDED")
	os.Setenv("RECPATH", "RECPATH")
	os.Setenv("LOGPATH", "LOGPATH")
	os.Setenv("ERROR_CNT", "4")
	os.Setenv("DROP_CNT", "5")
	os.Setenv("SCRAMBLING_CNT", "6")

	actual, err := LoadRecordingCommandEnv()
	if err != nil {
		t.Error(err)
	}

	expected := RecordingCommandEnv{
		RecordedID:           "1",
		ProgramID:            "2",
		ChannelType:          "CHANNELTYPE",
		ChannelID:            "3",
		ChannelName:          "CHANNELNAME",
		HalfWidthChannelName: "HALF_WIDTH_CHANNELNAME",
		StartAt:              "1715353200000",
		EndAt:                "1715355000000",
		Duration:             "1800000",
		Name:                 "NAME",
		HalfWidthName:        "HALF_WIDTH_NAME",
		Description:          "DESCRIPTION",
		HalfWidthDescription: "HALF_WIDTH_DESCRIPTION",
		Extended:             "EXTENDED",
		HalfWidthExtended:    "HALF_WIDTH_EXTENDED",
		RecPath:              "RECPATH",
		LogPath:              "LOGPATH",
		ErrorCnt:             "4",
		DropCnt:              "5",
		ScramblingCount:      "6",
	}

	if !reflect.DeepEqual(expected, actual) {
		t.Fatalf("expected = %v, actual = %v", expected, actual)
	}
}

func Test_LoadEncodingCommandEnv_デフォルトパラメータで読み込める(t *testing.T) {
	os.Clearenv()

	actual, err := LoadEncodingCommandEnv()
	if err != nil {
		t.Error(err)
	}

	expected := EncodingCommandEnv{
		RecordedID:           "",
		VideoFileID:          "",
		OutputPath:           "None",
		Mode:                 "None",
		ChannelID:            "",
		ChannelName:          "None",
		HalfWidthChannelName: "None",
		Name:                 "None",
		HalfWidthName:        "None",
		Description:          "None",
		HalfWidthDescription: "None",
		Extended:             "None",
		HalfWidthExtended:    "None",
	}

	if !reflect.DeepEqual(expected, actual) {
		t.Fatalf("expected = %v, actual = %v", expected, actual)
	}
}

func Test_LoadEncodingCommandEnv_フルパラメータで読み込める(t *testing.T) {
	os.Clearenv()
	os.Setenv("RECORDEDID", "1")
	os.Setenv("VIDEOFILEID", "2")
	os.Setenv("OUTPUTPATH", "OUTPUTPATH")
	os.Setenv("MODE", "MODE")
	os.Setenv("CHANNELID", "3")
	os.Setenv("CHANNELNAME", "CHANNELNAME")
	os.Setenv("HALF_WIDTH_CHANNELNAME", "HALF_WIDTH_CHANNELNAME")
	os.Setenv("NAME", "NAME")
	os.Setenv("HALF_WIDTH_NAME", "HALF_WIDTH_NAME")
	os.Setenv("DESCRIPTION", "DESCRIPTION")
	os.Setenv("HALF_WIDTH_DESCRIPTION", "HALF_WIDTH_DESCRIPTION")
	os.Setenv("EXTENDED", "EXTENDED")
	os.Setenv("HALF_WIDTH_EXTENDED", "HALF_WIDTH_EXTENDED")

	actual, err := LoadEncodingCommandEnv()
	if err != nil {
		t.Error(err)
	}

	expected := EncodingCommandEnv{
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
