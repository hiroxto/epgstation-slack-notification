package main

import (
	"github.com/slack-go/slack"
	"io"
	"log"
	"os"
	"path/filepath"
)

func createSlackClient(apiKey string, debug bool) *slack.Client {
	exeFilePath, err := os.Executable()
	if err != nil {
		log.Fatal(err.Error())
	}

	logfile, err := os.OpenFile(filepath.Join(filepath.Dir(exeFilePath), "slack.log"), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer logfile.Close()

	api := slack.New(apiKey, slack.OptionDebug(debug), slack.OptionLog(log.New(io.MultiWriter(logfile, os.Stdout), "slack-bot: ", log.Lshortfile|log.LstdFlags)))

	_, err = api.AuthTest()

	if err != nil {
		log.Fatal(err.Error())
	}

	return api
}

func getSlackChannel(globalChannel string, localChannel string) string {
	channel := globalChannel

	if len(localChannel) > 0 {
		channel = localChannel
	}

	return channel
}
