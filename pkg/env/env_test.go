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
		ProgramID:            "None",
		ChannelType:          "None",
		ChannelID:            "None",
		ChannelName:          "None",
		HalfWidthChannelName: "None",
		StartAt:              "None",
		EndAt:                "None",
		Duration:             "None",
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
	os.Setenv("PROGRAMID", "PROGRAMID")
	os.Setenv("CHANNELTYPE", "CHANNELTYPE")
	os.Setenv("CHANNELID", "CHANNELID")
	os.Setenv("CHANNELNAME", "CHANNELNAME")
	os.Setenv("HALF_WIDTH_CHANNELNAME", "HALF_WIDTH_CHANNELNAME")
	os.Setenv("STARTAT", "STARTAT")
	os.Setenv("ENDAT", "ENDAT")
	os.Setenv("DURATION", "DURATION")
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
		ProgramID:            "PROGRAMID",
		ChannelType:          "CHANNELTYPE",
		ChannelID:            "CHANNELID",
		ChannelName:          "CHANNELNAME",
		HalfWidthChannelName: "HALF_WIDTH_CHANNELNAME",
		StartAt:              "STARTAT",
		EndAt:                "ENDAT",
		Duration:             "DURATION",
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
		RecordedID:           "None",
		ProgramID:            "None",
		ChannelType:          "None",
		ChannelID:            "None",
		ChannelName:          "None",
		HalfWidthChannelName: "None",
		StartAt:              "None",
		EndAt:                "None",
		Duration:             "None",
		Name:                 "None",
		HalfWidthName:        "None",
		Description:          "None",
		HalfWidthDescription: "None",
		Extended:             "None",
		HalfWidthExtended:    "None",
		RecPath:              "None",
		LogPath:              "None",
		ErrorCnt:             0,
		DropCnt:              0,
		ScramblingCount:      0,
	}

	if !reflect.DeepEqual(expected, actual) {
		t.Fatalf("expected = %v, actual = %v", expected, actual)
	}
}

func Test_LoadRecordingCommandEnv_フルパラメータで読み込める(t *testing.T) {
	os.Clearenv()
	os.Setenv("RECORDEDID", "RECORDEDID")
	os.Setenv("PROGRAMID", "PROGRAMID")
	os.Setenv("CHANNELTYPE", "CHANNELTYPE")
	os.Setenv("CHANNELID", "CHANNELID")
	os.Setenv("CHANNELNAME", "CHANNELNAME")
	os.Setenv("HALF_WIDTH_CHANNELNAME", "HALF_WIDTH_CHANNELNAME")
	os.Setenv("STARTAT", "STARTAT")
	os.Setenv("ENDAT", "ENDAT")
	os.Setenv("DURATION", "DURATION")
	os.Setenv("NAME", "NAME")
	os.Setenv("HALF_WIDTH_NAME", "HALF_WIDTH_NAME")
	os.Setenv("DESCRIPTION", "DESCRIPTION")
	os.Setenv("HALF_WIDTH_DESCRIPTION", "HALF_WIDTH_DESCRIPTION")
	os.Setenv("EXTENDED", "EXTENDED")
	os.Setenv("HALF_WIDTH_EXTENDED", "HALF_WIDTH_EXTENDED")
	os.Setenv("RECPATH", "RECPATH")
	os.Setenv("LOGPATH", "LOGPATH")
	os.Setenv("ERROR_CNT", "1")
	os.Setenv("DROP_CNT", "2")
	os.Setenv("SCRAMBLING_CNT", "3")

	actual, err := LoadRecordingCommandEnv()
	if err != nil {
		t.Error(err)
	}

	expected := RecordingCommandEnv{
		RecordedID:           "RECORDEDID",
		ProgramID:            "PROGRAMID",
		ChannelType:          "CHANNELTYPE",
		ChannelID:            "CHANNELID",
		ChannelName:          "CHANNELNAME",
		HalfWidthChannelName: "HALF_WIDTH_CHANNELNAME",
		StartAt:              "STARTAT",
		EndAt:                "ENDAT",
		Duration:             "DURATION",
		Name:                 "NAME",
		HalfWidthName:        "HALF_WIDTH_NAME",
		Description:          "DESCRIPTION",
		HalfWidthDescription: "HALF_WIDTH_DESCRIPTION",
		Extended:             "EXTENDED",
		HalfWidthExtended:    "HALF_WIDTH_EXTENDED",
		RecPath:              "RECPATH",
		LogPath:              "LOGPATH",
		ErrorCnt:             1,
		DropCnt:              2,
		ScramblingCount:      3,
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
		RecordedID:           "None",
		VideoFileID:          "None",
		OutputPath:           "None",
		Mode:                 "None",
		ChannelID:            "None",
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
	os.Setenv("RECORDEDID", "RECORDEDID")
	os.Setenv("VIDEOFILEID", "VIDEOFILEID")
	os.Setenv("OUTPUTPATH", "OUTPUTPATH")
	os.Setenv("MODE", "MODE")
	os.Setenv("CHANNELID", "CHANNELID")
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

	if !reflect.DeepEqual(expected, actual) {
		t.Fatalf("expected = %v, actual = %v", expected, actual)
	}
}
