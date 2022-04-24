package main

import (
	"bytes"
	"fmt"
	"github.com/slack-go/slack"
	"text/template"
)

func buildCommandFields(fieldsConfigs []FieldConfig, env interface{}) ([]*slack.TextBlockObject, error) {
	var fields []*slack.TextBlockObject

	for _, fieldsConfig := range fieldsConfigs {
		content, err := formatCommandEnv("", fieldsConfig.Template, env)

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

func formatCommandEnv(name string, userTemplate string, env interface{}) (string, error) {
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

func buildRecordedLogReportFields(recordedLog RecordedLog) ([]*slack.TextBlockObject, error) {
	var fields []*slack.TextBlockObject

	fields = append(fields, createNewTextBlockField("RecordedID", fmt.Sprint(recordedLog.ID)))
	fields = append(fields, createNewTextBlockField("ErrorCnt", fmt.Sprint(recordedLog.ErrorCnt)))
	fields = append(fields, createNewTextBlockField("DropCnt", fmt.Sprint(recordedLog.DropCnt)))
	fields = append(fields, createNewTextBlockField("ScramblingCnt", fmt.Sprint(recordedLog.ScramblingCnt)))

	return fields, nil
}
