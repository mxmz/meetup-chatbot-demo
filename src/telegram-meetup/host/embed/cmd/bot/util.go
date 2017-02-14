package main

import (
	"encoding/json"
	"io"
	"io/ioutil"
)

func ReadJson(reader io.ReadCloser, obj interface{}) error {
	data, err := ioutil.ReadAll(reader)
	if err != nil {
		return err
	} else {
		err := json.Unmarshal(data, obj)
		if err != nil {
			return err
		} else {
			return nil
		}
	}
}

func nvl(ss ...string) string {
	for _, s := range ss {
		if len(s) > 0 {
			return s
		}
	}
	return ""
}
