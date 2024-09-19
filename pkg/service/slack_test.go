package service

import "testing"

func Test_formatContent(t *testing.T) {
	userTemplate := "{{ .Foo }}-test"
	detail := struct {
		Foo string
	}{
		Foo: "Bar",
	}

	expected := "Bar-test"
	actual, err := formatContent("dummy", userTemplate, detail)

	if err != nil {
		t.Errorf("formatContent() returned an unexpected error: %v", err)
	}

	if actual != expected {
		t.Errorf("formatContent() = %v, want %v", actual, expected)
	}
}
