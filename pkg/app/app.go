package app

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"text/template"

	"github.com/hiroxto/epgstation-slack-notification/pkg/config"
	"github.com/hiroxto/epgstation-slack-notification/pkg/service"
	"github.com/slack-go/slack"
)

// Field 通知情報のセクションの中身
type Field struct {
	Title    string
	Template string
}

// FieldsFromConfig config.FieldConfigをFieldに変換する
func FieldsFromConfig(fieldConfigs []config.FieldConfig) []Field {
	var fields []Field

	for _, fieldConfig := range fieldConfigs {
		fields = append(fields, Field{
			Title:    fieldConfig.Title,
			Template: fieldConfig.Template,
		})
	}

	return fields
}

// convertFields app.Fieldsをservice.Fieldに変換する
func convertFields(appFields []Field) []service.Field {
	var fields []service.Field

	for _, fieldConfig := range appFields {
		fields = append(fields, service.Field{
			Title:    fieldConfig.Title,
			Template: fieldConfig.Template,
		})
	}

	return fields
}

// createSlackClient Slackクライアントを作成する
func createSlackClient(apiKey string, debug bool) *slack.Client {
	return slack.New(apiKey, slack.OptionDebug(debug), slack.OptionLog(log.New(os.Stdout, "slack-bot: ", log.Lshortfile|log.LstdFlags)))
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

// buildCommandFields テンプレートをフォーマットしたフィールドのリストを作る
func buildCommandFields(fieldsConfigs []Field, detail interface{}) ([]*slack.TextBlockObject, error) {
	var fields []*slack.TextBlockObject

	for _, fieldsConfig := range fieldsConfigs {
		content, err := formatContent("", fieldsConfig.Template, detail)

		if err != nil {
			return nil, err
		}

		text := fmt.Sprintf("*%s*\n%s", fieldsConfig.Title, content)
		fields = append(fields, slack.NewTextBlockObject(slack.MarkdownType, text, false, false))
	}

	return fields, nil
}

// buildMessageOptions メッセージオプションを作成する
func buildMessageOptions(message string, fields []*slack.TextBlockObject) slack.MsgOption {
	fallbackOpt := slack.MsgOptionText(message, false)
	blockOpt := slack.MsgOptionBlocks(
		createHeaderSection(message),
		createFieldsSection(fields),
		slack.NewDividerBlock(),
	)

	return slack.MsgOptionCompose(fallbackOpt, blockOpt)
}

// createHeaderSection header部分を作成する
func createHeaderSection(text string) *slack.SectionBlock {
	headerText := slack.NewTextBlockObject(slack.MarkdownType, text, false, false)

	return slack.NewSectionBlock(headerText, nil, nil)
}

// createFieldsSection 情報ブロックを作成する
func createFieldsSection(fields []*slack.TextBlockObject) *slack.SectionBlock {
	return slack.NewSectionBlock(nil, fields, nil)
}
