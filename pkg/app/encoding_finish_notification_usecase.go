package app

// EncodingFinishNotificationParam EncodingFinishNotificationUseCaseのパラメータ
type EncodingFinishNotificationParam struct {
	EnableDebug          bool
	SlackAPIKey          string
	SlackChannel         string
	Message              string
	Fields               []Field
	EncodingFinishDetail EncodingFinishDetail
}

// EncodingFinishDetail エンコーディング終了情報
type EncodingFinishDetail struct {
	RecordedID  string
	VideoFileID string
	OutputPath  string
	Mode        string
	ChannelID   string
	ChannelName string
	Name        string
	Description string
	Extended    string
	Original    EncodingFinishOriginal
}

// EncodingFinishOriginal エンコーディング終了情報の環境変数そのままの値
type EncodingFinishOriginal struct {
	RecordedID           string
	VideoFileID          string
	OutputPath           string
	Mode                 string
	ChannelID            string
	ChannelName          string
	HalfWidthChannelName string
	Name                 string
	HalfWidthName        string
	Description          string
	HalfWidthDescription string
	Extended             string
	HalfWidthExtended    string
}

// EncodingFinishNotificationUseCase エンコーディングの終了を通知する
func EncodingFinishNotificationUseCase(param EncodingFinishNotificationParam) error {
	slackClient, err := createSlackClient(param.SlackAPIKey, param.EnableDebug)
	if err != nil {
		return err
	}

	message, err := formatContent("", param.Message, param.EncodingFinishDetail)
	if err != nil {
		return err
	}

	fields, err := buildCommandFields(param.Fields, param.EncodingFinishDetail)
	if err != nil {
		return err
	}

	options := buildMessageOptions(message, fields)
	_, _, err = slackClient.PostMessage(param.SlackChannel, options)

	if err != nil {
		return err
	}

	return nil
}
