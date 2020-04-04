package main

import (
	"bytes"
	"fmt"
	"github.com/slack-go/slack"
	"text/template"
)

func buildPreCommandFields(fieldsConfigs []FieldConfig, env PreCommandEnv) ([]*slack.TextBlockObject, error) {
	var fields []*slack.TextBlockObject

	for _, fieldsConfig := range fieldsConfigs {
		content, err := formatPreCommandEnv("", fieldsConfig.Template, env)

		if err != nil {
			return nil, err
		}

		fields = append(fields, createNewTextBlockField(fieldsConfig.Title, content))
	}

	return fields, nil
}

func buildRecCommandFields(fieldsConfigs []FieldConfig, env RecCommandEnv) ([]*slack.TextBlockObject, error) {
	var fields []*slack.TextBlockObject

	for _, fieldsConfig := range fieldsConfigs {
		content, err := formatRecCommandEnv("", fieldsConfig.Template, env)

		if err != nil {
			return nil, err
		}

		fields = append(fields, createNewTextBlockField(fieldsConfig.Title, content))
	}

	return fields, nil
}

func createNewTextBlockField(title string, body string) *slack.TextBlockObject {
	text := fmt.Sprintf("*%s*\n%s", title, body)
	return slack.NewTextBlockObject("mrkdwn", text, false, false)
}

func formatCommandEnv(name string, userTemplate string, env CommandEnv) (string, error) {
	var messageBuffer bytes.Buffer
	t, err := template.New(name).Parse(userTemplate)

	if err != nil {
		return messageBuffer.String(), err
	}

	if err := t.Execute(&messageBuffer, env); err != nil {
		return messageBuffer.String(), err
	}

	return messageBuffer.String(), nil
}

func formatPreCommandEnv(name string, userTemplate string, env PreCommandEnv) (string, error) {
	var messageBuffer bytes.Buffer
	t, err := template.New(name).Parse(userTemplate)

	if err != nil {
		return messageBuffer.String(), err
	}

	if err := t.Execute(&messageBuffer, env); err != nil {
		return messageBuffer.String(), err
	}

	return messageBuffer.String(), nil
}

func formatRecCommandEnv(name string, userTemplate string, env RecCommandEnv) (string, error) {
	var messageBuffer bytes.Buffer
	t, err := template.New(name).Parse(userTemplate)

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

func buildMessageOptions(message string, fields []*slack.TextBlockObject) slack.MsgOption {
	fallbackOpt := slack.MsgOptionText(message, false)
	blockOpt := slack.MsgOptionBlocks(
		createHeaderSection(message),
		createFieldsSection(fields),
		slack.NewDividerBlock(),
	)

	return slack.MsgOptionCompose(fallbackOpt, blockOpt)
}
