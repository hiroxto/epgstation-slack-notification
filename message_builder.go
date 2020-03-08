package main

import (
	"fmt"
	"github.com/slack-go/slack"
	"strings"
)

func buildPreCommandBlocks(message string, env PreCommandEnv) slack.MsgOption {
	headerText := slack.NewTextBlockObject("mrkdwn", message, false, false)
	headerSection := slack.NewSectionBlock(headerText, nil, nil)

	channels := []string{
		env.ChannelType,
		env.ChannelID,
		env.ChannelName,
	}
	times := []string{
		env.StartAt,
		env.EndAt,
	}

	fields := []*slack.TextBlockObject{
		buildTextBlock("ProgramID", env.ProgramID),
		buildTextBlock("ChannelType, ChannelID, ChannelName", strings.Join(channels, "\n")),
		buildTextBlock("StartAt, EndAt", strings.Join(times, "\n")),
		buildTextBlock("Duration", env.Duration),
		buildTextBlock("Name", env.Name),
		buildTextBlock("Description", env.Description),
		buildTextBlock("Extended", env.Extended),
	}
	fieldsSection := slack.NewSectionBlock(nil, fields, nil)

	fallbackOpt := slack.MsgOptionText(message, false)
	blockOpt := slack.MsgOptionBlocks(
		headerSection,
		fieldsSection,
		slack.NewDividerBlock(),
	)

	return slack.MsgOptionCompose(fallbackOpt, blockOpt)
}

func buildRecCommandBlocks(message string, env RecCommandEnv) slack.MsgOption {
	headerText := slack.NewTextBlockObject("mrkdwn", message, false, false)
	headerSection := slack.NewSectionBlock(headerText, nil, nil)

	channels := []string{
		env.ChannelType,
		env.ChannelID,
		env.ChannelName,
	}
	times := []string{
		env.StartAt,
		env.EndAt,
	}

	fields := []*slack.TextBlockObject{
		buildTextBlock("RecordedID", env.RecordedID),
		buildTextBlock("ProgramID", env.ProgramID),
		buildTextBlock("ChannelType, ChannelID, ChannelName", strings.Join(channels, "\n")),
		buildTextBlock("StartAt, EndAt", strings.Join(times, "\n")),
		buildTextBlock("Duration", env.Duration),
		buildTextBlock("Name", env.Name),
		buildTextBlock("Description", env.Description),
		buildTextBlock("Extended", env.Extended),
		buildTextBlock("RecPath", env.RecPath),
		buildTextBlock("LogPath", env.LogPath),
	}
	fieldsSection := slack.NewSectionBlock(nil, fields, nil)

	fallbackOpt := slack.MsgOptionText(message, false)
	blockOpt := slack.MsgOptionBlocks(
		headerSection,
		fieldsSection,
		slack.NewDividerBlock(),
	)

	return slack.MsgOptionCompose(fallbackOpt, blockOpt)
}

func buildTextBlock(title string, value string) *slack.TextBlockObject {
	text := fmt.Sprintf("*%s:*\n%s", title, value)

	return slack.NewTextBlockObject("mrkdwn", text, false, false)
}
