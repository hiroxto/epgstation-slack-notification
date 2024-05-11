package command

import (
	"reflect"
	"testing"
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
