package service

import (
	"bytes"
	"fmt"
	"text/template"

	"github.com/slack-go/slack"
)

type SlackService struct {
	client *slack.Client
}

// Field 通知情報のセクションの中身
type Field struct {
	Title    string
	Template string
}

func NewSlackService(apiKey string, debug bool) *SlackService {
	client := slack.New(apiKey, slack.OptionDebug(debug))
	return &SlackService{client: client}
}

// PostMessageWithFields メッセージとFieldをフォーマットして投稿する
func (c *SlackService) PostMessageWithFields(channelID string, messageTemplate string, fields []Field, detail interface{}, userName string) error {
	message, err := formatContent("message", messageTemplate, detail)
	if err != nil {
		return err
	}

	var textBlocks []*slack.TextBlockObject
	for i, field := range fields {
		content, err := formatContent(fmt.Sprintf("field:%d", i), field.Template, detail)

		if err != nil {
			return err
		}

		text := fmt.Sprintf("*%s*\n%s", field.Title, content)
		textBlocks = append(textBlocks, slack.NewTextBlockObject(slack.MarkdownType, text, false, false))
	}

	fallbackOption := slack.MsgOptionText(message, false)
	blockOption := slack.MsgOptionBlocks(
		slack.NewSectionBlock(
			slack.NewTextBlockObject(slack.MarkdownType, message, false, false),
			nil,
			nil,
		),
		slack.NewSectionBlock(
			nil,
			textBlocks,
			nil,
		),
	)

	messageOptions := slack.MsgOptionCompose(fallbackOption, blockOption)
	userNameOption := slack.MsgOptionUsername(userName)
	_, _, err = c.client.PostMessage(channelID, messageOptions, userNameOption)
	if err != nil {
		return err
	}

	return nil
}

// formatContent テンプレートをフォーマットする
func formatContent(name string, userTemplate string, detail interface{}) (string, error) {
	var messageBuffer bytes.Buffer
	t, err := template.New(name).Parse(userTemplate)

	if err != nil {
		return messageBuffer.String(), err
	}

	if err := t.Execute(&messageBuffer, detail); err != nil {
		return messageBuffer.String(), err
	}

	return messageBuffer.String(), nil
}
