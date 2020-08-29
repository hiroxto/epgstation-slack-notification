package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

// RecordedLog 録画情報を入れるstruct
type RecordedLog struct {
	ID            int `json:"id"`
	ErrorCnt      int `json:"errorCnt"`
	DropCnt       int `json:"dropCnt"`
	ScramblingCnt int `json:"scramblingCnt"`
}

func getRecordedLog(id string, config Config) (RecordedLog, error) {
	apiResponse, err := callRecordedAPI(config.EPGStation.HostName, id)
	if err != nil {
		return RecordedLog{}, err
	}
	recordedLog, err := jsonBytesToRecordedLog(apiResponse)
	if err != nil {
		return RecordedLog{}, err
	}

	return recordedLog, nil
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
