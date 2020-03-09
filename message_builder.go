package main

import (
	"bytes"
	"fmt"
	"github.com/slack-go/slack"
	"strings"
	"text/template"
)

func buildPreCommandHeaderText(messageTemplate string, env PreCommandEnv) (string, error) {
	var messageBuffer bytes.Buffer
	t, err := template.New("message").Parse(messageTemplate)

	if err != nil {
		return messageBuffer.String(), err
	}

	if err := t.Execute(&messageBuffer, env); err != nil {
		return messageBuffer.String(), err
	}

	return messageBuffer.String(), nil
}

func buildRecCommandHeaderText(messageTemplate string, env RecCommandEnv) (string, error) {
	var messageBuffer bytes.Buffer
	t, err := template.New("message").Parse(messageTemplate)

	if err != nil {
		return messageBuffer.String(), err
	}

	if err := t.Execute(&messageBuffer, env); err != nil {
		return messageBuffer.String(), err
	}

	return messageBuffer.String(), nil
}

func createHeaderSection(text string) *slack.SectionBlock {
	headerText := slack.NewTextBlockObject("mrkdwn", text, false, false)

	return slack.NewSectionBlock(headerText, nil, nil)
}

func createFieldsSection(fields []*slack.TextBlockObject) *slack.SectionBlock {
	return slack.NewSectionBlock(nil, fields, nil)
}

func buildPreCommandMessageOptions(message string, env PreCommandEnv) slack.MsgOption {
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

	fallbackOpt := slack.MsgOptionText(message, false)
	blockOpt := slack.MsgOptionBlocks(
		createHeaderSection(message),
		createFieldsSection(fields),
		slack.NewDividerBlock(),
	)

	return slack.MsgOptionCompose(fallbackOpt, blockOpt)
}

func buildRecCommandBlocks(message string, env RecCommandEnv) slack.MsgOption {
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

	fallbackOpt := slack.MsgOptionText(message, false)
	blockOpt := slack.MsgOptionBlocks(
		createHeaderSection(message),
		createFieldsSection(fields),
		slack.NewDividerBlock(),
	)

	return slack.MsgOptionCompose(fallbackOpt, blockOpt)
}

func buildTextBlock(title string, value string) *slack.TextBlockObject {
	text := fmt.Sprintf("*%s:*\n%s", title, value)

	return slack.NewTextBlockObject("mrkdwn", text, false, false)
}
