package main

import (
	"github.com/slack-go/slack"
	"log"
	"os"
)

func createSlackClient(apiKey string, debug bool) (*slack.Client, error) {
	api := slack.New(apiKey, slack.OptionDebug(debug), slack.OptionLog(log.New(os.Stdout, "slack-bot: ", log.Lshortfile|log.LstdFlags)))

	if _, err := api.AuthTest(); err != nil {
		return nil, err
	}

	return api, nil
}
