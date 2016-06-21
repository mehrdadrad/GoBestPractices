package config

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

type Config struct {
	Root interface{}
}

func (cfg *Config) LoadJson(filename string) map[string]interface{} {
	content, err := ioutil.ReadFile(filename)
	var out map[string]interface{}
	err = json.Unmarshal([]byte(content), &out)
	if err != nil {
		log.Fatal(err)
	}
	return out
}
