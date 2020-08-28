package main

import (
	"io/ioutil"
	"net/http"
)

func callRecordedAPI(hostName string, id string) (string, error) {
	resp, err := http.Get(hostName + "/api/recorded/" + id)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(bytes), nil
}
