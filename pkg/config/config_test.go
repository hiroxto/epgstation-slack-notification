package config

import (
	"reflect"
	"testing"
)

func Test_LoadConfigFromYaml_YAMLから読み込める(t *testing.T) {
	configYaml := `
epg-station:
    host-name: "http://localhost:8888"

slack:
    api-key: "YOUR_API_KEY"
    channel: "CHANNEL_ID"

commands:
    reserve-new-addition: &pre-command-default
        enable: true
        message: ":new: {{ .ChannelName }} で {{ .Name }} の録画予約が新規追加されました"
        fields-section:
            -   title: "ProgramID"
                template: "{{ .ProgramID }}"
        channel: "OVERRIDE_CHANNEL_ID"

    reserve-update:
        <<: *pre-command-default
        enable: true
        message: ":up: {{ .ChannelName }} で {{ .Name }} の録画情報が更新されました"

    reserve-deleted:
        <<: *pre-command-default
        enable: true
        message: ":black_square_for_stop: {{ .ChannelName }} で {{ .Name }} の録画予約が削除されました"

    recording-pre-start:
        <<: *pre-command-default
        enable: true
        message: ":soon: {{ .ChannelName }} で {{ .Name }} の録画準備が開始しました"

    recording-prep-rec-failed:
        <<: *pre-command-default
        enable: true
        message: ":x: {{ .ChannelName }} で {{ .Name }} の録画準備に失敗しました"

    recording-start: &rec-command-default
        enable: true
        message: ":arrow_forward: {{ .ChannelName }} で {{ .Name }} の録画が開始しました"
        fields-section:
            -   title: "RecordedID, ProgramID"
                template: "{{ .RecordedID }}, {{ .ProgramID }}"

    recording-finish:
        <<: *rec-command-default
        enable: true
        message: ":white_check_mark: {{ .ChannelName }} で {{ .Name }} の録画が終了しました"

    recording-failed:
        <<: *rec-command-default
        enable: true
        message: ":x: {{ .ChannelName }} で {{ .Name }} の録画中にエラーが発生しました"

    encoding-finish:
        enable: true
        message: ":white_check_mark: {{ .HalfWidthChannelName }} の {{ .HalfWidthName }} のエンコードが終了しました"
        fields-section:
            -   title: "RecordedID, VideoFileID"
                template: "{{ .RecordedID }}, {{ .VideoFileID }}"
`

	actual, err := LoadConfigFromYaml([]byte(configYaml))
	if err != nil {
		t.Error(err)
	}

	expected := Config{
		EPGStation: EPGStationConfig{
			HostName: "http://localhost:8888",
		},
		Slack: SlackConfig{
			APIKey:  "YOUR_API_KEY",
			Channel: "CHANNEL_ID",
		},
		Commands: CommandsConfig{
			ReserveNewAddition: CommandConfig{
				Enable:  true,
				Channel: "OVERRIDE_CHANNEL_ID",
				Message: ":new: {{ .ChannelName }} で {{ .Name }} の録画予約が新規追加されました",
				FieldsSection: []FieldConfig{
					{
						Title:    "ProgramID",
						Template: "{{ .ProgramID }}",
					},
				},
			},
			ReserveUpdateCommand: CommandConfig{
				Enable:  true,
				Channel: "OVERRIDE_CHANNEL_ID",
				Message: ":up: {{ .ChannelName }} で {{ .Name }} の録画情報が更新されました",
				FieldsSection: []FieldConfig{
					{
						Title:    "ProgramID",
						Template: "{{ .ProgramID }}",
					},
				},
			},
			ReserveDeletedCommand: CommandConfig{
				Enable:  true,
				Channel: "OVERRIDE_CHANNEL_ID",
				Message: ":black_square_for_stop: {{ .ChannelName }} で {{ .Name }} の録画予約が削除されました",
				FieldsSection: []FieldConfig{
					{
						Title:    "ProgramID",
						Template: "{{ .ProgramID }}",
					},
				},
			},
			RecordingPreStart: CommandConfig{
				Enable:  true,
				Channel: "OVERRIDE_CHANNEL_ID",
				Message: ":soon: {{ .ChannelName }} で {{ .Name }} の録画準備が開始しました",
				FieldsSection: []FieldConfig{
					{
						Title:    "ProgramID",
						Template: "{{ .ProgramID }}",
					},
				},
			},
			RecordingPrepRecFailed: CommandConfig{
				Enable:  true,
				Channel: "OVERRIDE_CHANNEL_ID",
				Message: ":x: {{ .ChannelName }} で {{ .Name }} の録画準備に失敗しました",
				FieldsSection: []FieldConfig{
					{
						Title:    "ProgramID",
						Template: "{{ .ProgramID }}",
					},
				},
			},
			RecordingStart: CommandConfig{
				Enable:  true,
				Message: ":arrow_forward: {{ .ChannelName }} で {{ .Name }} の録画が開始しました",
				FieldsSection: []FieldConfig{
					{
						Title:    "RecordedID, ProgramID",
						Template: "{{ .RecordedID }}, {{ .ProgramID }}",
					},
				},
			},
			RecordingFinish: CommandConfig{
				Enable:  true,
				Message: ":white_check_mark: {{ .ChannelName }} で {{ .Name }} の録画が終了しました",
				FieldsSection: []FieldConfig{
					{
						Title:    "RecordedID, ProgramID",
						Template: "{{ .RecordedID }}, {{ .ProgramID }}",
					},
				},
			},
			RecordingFailed: CommandConfig{
				Enable:  true,
				Message: ":x: {{ .ChannelName }} で {{ .Name }} の録画中にエラーが発生しました",
				FieldsSection: []FieldConfig{
					{
						Title:    "RecordedID, ProgramID",
						Template: "{{ .RecordedID }}, {{ .ProgramID }}",
					},
				},
			},
			EncodingFinish: CommandConfig{
				Enable:  true,
				Message: ":white_check_mark: {{ .HalfWidthChannelName }} の {{ .HalfWidthName }} のエンコードが終了しました",
				FieldsSection: []FieldConfig{
					{
						Title:    "RecordedID, VideoFileID",
						Template: "{{ .RecordedID }}, {{ .VideoFileID }}",
					},
				},
			},
		},
	}

	if !reflect.DeepEqual(expected, actual) {
		t.Fatalf("expected = %v, actual = %v", expected, actual)
	}
}
