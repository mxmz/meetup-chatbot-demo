package env

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

var DevMode = false

type EnvMap map[string]string

func (e EnvMap) get(key string) string {
	s := e[key]
	if len(s) == 0 {
		return os.Getenv(key)
	} else {
		return s
	}
}

func (e EnvMap) Get(key string) string {
	if DevMode {
		s := e.get("dev_" + key)
		if len(s) > 0 {
			return s
		}
	}
	return e.get(key)
}

func LoadEnv() EnvMap {
	var em EnvMap
	envfile, err := ioutil.ReadFile("./env.json")
	if err != nil {
		panic(err.Error())
	}
	json.Unmarshal(envfile, &em)
	return em
}

func (e *EnvMap) Export() {
	for k, v := range *e {
		os.Setenv(k, v)
	}
}
