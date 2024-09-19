package app

import (
	"github.com/hiroxto/epgstation-slack-notification/pkg/config"
	"github.com/hiroxto/epgstation-slack-notification/pkg/service"
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
