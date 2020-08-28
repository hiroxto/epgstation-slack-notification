package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type RecordedLog struct {
	ID            string
	ErrorCnt      int
	DropCnt       int
	ScramblingCnt int
}

func callRecordedAPI(hostName string, id string) ([]byte, error) {
	resp, err := http.Get(hostName + "/api/recorded/" + id)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return bytes, nil
}

func jsonBytesToRecordedLog(data []byte) (RecordedLog, error) {
	var log RecordedLog
	if err := json.Unmarshal(data, &log); err != nil {
		return RecordedLog{}, err
	}

	return log, nil
}