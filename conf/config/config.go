package config

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

type Config struct {
	Root map[string]interface{}
}

func (cfg *Config) LoadJSON(filename string) {
	content, err := ioutil.ReadFile(filename)
	var out map[string]interface{}
	err = json.Unmarshal([]byte(content), &out)
	if err != nil {
		log.Fatal(err)
	}
	cfg.Root = out
}

func (cfg *Config) Get(keys string) interface{} {
	return cfg.Root[keys]
}
