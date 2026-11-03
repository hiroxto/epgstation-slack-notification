package main

import (
	"bytes"
	"context"
	"io"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/k0kubun/pp/v3"
)

func TestCLIVersion(t *testing.T) {
	originalVersion, originalCommit, originalDate := version, commit, date
	t.Cleanup(func() { version, commit, date = originalVersion, originalCommit, originalDate })
	version = "2.2.2"
	commit = "0123456789abcdef0123456789abcdef01234567"
	date = "2026-11-03T03:00:00Z"

	app, err := newCliApp()
	if err != nil {
		t.Fatal(err)
	}
	var output bytes.Buffer
	app.Writer = &output
	if err := app.Run(context.Background(), []string{"notification", "--version"}); err != nil {
		t.Fatal(err)
	}
	want := "epgstation-slack-notification version 2.2.2 (commit: 0123456789abcdef0123456789abcdef01234567, date: 2026-11-03T03:00:00Z)\n"
	if output.String() != want {
		t.Fatalf("unexpected version output: got %q, want %q", output.String(), want)
	}
}

func TestCLIConfigFlags(t *testing.T) {
	configPath := filepath.Join(t.TempDir(), "disabled.yml")
	if err := os.WriteFile(configPath, []byte("commands: {}\n"), 0600); err != nil {
		t.Fatal(err)
	}

	for _, args := range [][]string{
		{"--config", configPath, "recording-start"},
		{"-c", configPath, "recorded-start"},
		{"recording-start", "--config", configPath},
	} {
		t.Run(strings.Join(args, " "), func(t *testing.T) {
			app, err := newCliApp()
			if err != nil {
				t.Fatal(err)
			}
			output := captureStdout(t, func() {
				if err := app.Run(context.Background(), append([]string{"notification"}, args...)); err != nil {
					t.Fatal(err)
				}
			})
			if output != "recording-start command is disabled.\n" {
				t.Fatalf("unexpected output: %q", output)
			}
		})
	}
}

func TestCLIMissingConfig(t *testing.T) {
	configPath := filepath.Join(t.TempDir(), "missing.yml")
	app, err := newCliApp()
	if err != nil {
		t.Fatal(err)
	}
	err = app.Run(context.Background(), []string{"notification", "-c", configPath, "recording-start"})
	if !os.IsNotExist(err) {
		t.Fatalf("expected missing config error, got %v", err)
	}
}

func TestCLIDumpEnv(t *testing.T) {
	t.Setenv("PROGRAMID", "migration-test-program")
	app, err := newCliApp()
	if err != nil {
		t.Fatal(err)
	}
	var buffer bytes.Buffer
	original := pp.GetDefaultOutput()
	pp.SetDefaultOutput(&buffer)
	t.Cleanup(func() { pp.SetDefaultOutput(original) })
	if err := app.Run(context.Background(), []string{"notification", "dump:env", "--color=false"}); err != nil {
		t.Fatal(err)
	}
	output := buffer.String()
	if !strings.Contains(output, "migration-test-program") {
		t.Fatalf("environment value missing from output: %q", output)
	}
	if strings.Contains(output, "\x1b[") {
		t.Fatalf("color disabled but output contains ANSI escape sequence: %q", output)
	}
}

func captureStdout(t *testing.T, action func()) string {
	t.Helper()
	reader, writer, err := os.Pipe()
	if err != nil {
		t.Fatal(err)
	}
	original := os.Stdout
	os.Stdout = writer
	defer func() {
		os.Stdout = original
		reader.Close()
		writer.Close()
	}()
	action()
	writer.Close()
	output, err := io.ReadAll(reader)
	if err != nil {
		t.Fatal(err)
	}
	return string(output)
}
