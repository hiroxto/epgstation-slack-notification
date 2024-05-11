package main

import (
	"github.com/hiroxto/epgstation-slack-notification/pkg/command"
	"github.com/urfave/cli/v2"
)

var commands = []*cli.Command{
	command.ReserveNewAdditionCommand,
	command.ReserveUpdateCommand,
	command.ReserveDeletedCommand,
	command.RecordingPreStartCommand,
	command.RecordingPrepRecFailedCommand,

	command.RecordingStartCommand,
	command.RecordingFinishCommand,
	command.RecordingFailedCommand,

	command.EncodingFinishCommand,

	command.DumpEnvsCommand,
}
